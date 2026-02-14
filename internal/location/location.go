package location

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Location represents a resolved geographic position.
type Location struct {
	Latitude  float64
	Longitude float64
	City      string
	Country   string
	Source    string // "corelocation", "ip", "manual"
}

// Config holds runtime configuration from CLI flags.
type Config struct {
	City      string
	Latitude  float64
	Longitude float64
	Imperial  bool
	NoColor   bool
	Days      int
}

// GeocodeFunc is a function type for city-to-location geocoding.
// Set by the weather package to avoid circular imports.
var GeocodeFunc func(city string) (float64, float64, string, string, error)

// ResolveLocation determines the user's location based on config.
// Priority: lat/lon flags > city flag > CoreLocation > IP geolocation.
func ResolveLocation(cfg Config) (Location, error) {
	if cfg.Latitude != 0 || cfg.Longitude != 0 {
		return Location{
			Latitude:  cfg.Latitude,
			Longitude: cfg.Longitude,
			Source:    "manual",
		}, nil
	}

	if cfg.City != "" && GeocodeFunc != nil {
		lat, lon, city, country, err := GeocodeFunc(cfg.City)
		if err != nil {
			return Location{}, err
		}
		return Location{
			Latitude:  lat,
			Longitude: lon,
			City:      city,
			Country:   country,
			Source:    "manual",
		}, nil
	}

	// Try CoreLocation first, fall back to IP silently
	loc, err := GetCoreLocation()
	if err == nil {
		return loc, nil
	}

	return getIPLocation()
}

// getIPLocation uses ip-api.com to determine location from IP address.
func getIPLocation() (Location, error) {
	return fetchIPLocation("http://ip-api.com/json/", &http.Client{Timeout: 5 * time.Second})
}

func fetchIPLocation(url string, client *http.Client) (Location, error) {
	resp, err := client.Get(url)
	if err != nil {
		return Location{}, fmt.Errorf("IP geolocation request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("IP geolocation returned status %d", resp.StatusCode)
	}

	var result struct {
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
		City        string  `json:"city"`
		RegionName  string  `json:"regionName"`
		Country     string  `json:"country"`
		Status      string  `json:"status"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Location{}, fmt.Errorf("failed to parse IP geolocation response: %w", err)
	}

	if result.Status != "success" {
		return Location{}, fmt.Errorf("IP geolocation failed")
	}

	city := result.City
	if city == "" {
		city = result.RegionName
	}

	return Location{
		Latitude:  result.Lat,
		Longitude: result.Lon,
		City:      city,
		Country:   result.Country,
		Source:    "ip",
	}, nil
}
