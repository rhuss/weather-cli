package main

import (
	"flag"
	"fmt"
	"goweather/internal/display"
	"goweather/internal/i18n"
	"goweather/internal/location"
	"goweather/internal/weather"
	"os"
)

func main() {
	city := flag.String("city", "", "City name for weather lookup")
	lat := flag.Float64("lat", 0, "Latitude for weather lookup")
	lon := flag.Float64("lon", 0, "Longitude for weather lookup")
	imperial := flag.Bool("imperial", false, "Use imperial units (Fahrenheit, mph)")
	metric := flag.Bool("metric", false, "Use metric units (Celsius, km/h) [default]")
	noColor := flag.Bool("no-color", false, "Disable ANSI color codes in output")
	days := flag.Int("days", 5, "Number of forecast days (1-7)")
	lang := flag.String("lang", "", "Language (en, de, es, fr, it, zh)")
	flag.Parse()

	// Initialize i18n (before any output)
	i18n.Init(*lang)

	// metric is the default; --metric is a no-op but accepted for clarity
	_ = metric

	// Validate --days range
	if *days < 1 || *days > 7 {
		fmt.Fprintf(os.Stderr, "Error: --days must be between 1 and 7 (got %d)\n", *days)
		os.Exit(1)
	}

	// Validate --lat/--lon pairing
	if (*lat != 0 && *lon == 0) || (*lat == 0 && *lon != 0) {
		fmt.Fprintln(os.Stderr, "Error: Both --lat and --lon must be provided together")
		os.Exit(1)
	}

	// Apply color setting
	display.ColorEnabled = !*noColor

	// Wire up geocoding function to avoid circular imports
	location.GeocodeFunc = weather.GeocodeCity

	cfg := location.Config{
		City:      *city,
		Latitude:  *lat,
		Longitude: *lon,
		Imperial:  *imperial,
		NoColor:   *noColor,
		Days:      *days,
	}

	// Resolve location
	loc, err := location.ResolveLocation(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if loc.Source == "" {
			fmt.Fprintln(os.Stderr, i18n.TipManualLocation())
		}
		os.Exit(1)
	}

	// Fetch weather
	client := weather.NewClient()
	data, err := client.FetchWeather(loc.Latitude, loc.Longitude, cfg.Days, cfg.Imperial)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Unable to fetch weather data: %v\n", err)
		os.Exit(1)
	}

	// Build location display name
	locName := loc.City
	if loc.Country != "" {
		if locName != "" {
			locName += ", " + loc.Country
		} else {
			locName = loc.Country
		}
	}
	if locName == "" {
		locName = fmt.Sprintf("%.2f, %.2f", loc.Latitude, loc.Longitude)
	}

	// Render and print
	output := display.RenderWeatherCard(locName, data, cfg.Imperial, cfg.Days)
	fmt.Print(output)
}
