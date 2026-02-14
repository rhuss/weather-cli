package display

import (
	"fmt"
	"goweather/internal/units"
	"goweather/internal/weather"
	"strings"
	"time"
)

const cardWidth = 42

// RenderWeatherCard produces the full terminal output for weather data.
func RenderWeatherCard(loc string, data *weather.WeatherData, imperial bool, days int) string {
	var b strings.Builder

	cond := GetCondition(data.Current.WeatherCode)
	art := AsciiArt(cond.Category)
	artLines := strings.Split(art, "\n")

	// Top border
	b.WriteString(topBorder())

	// Header: location + emoji
	header := fmt.Sprintf("  %s  %s", Bold(loc), cond.Emoji)
	b.WriteString(padLine(header))

	b.WriteString(emptyLine())

	// Current conditions alongside ASCII art
	infoLines := []string{
		fmt.Sprintf("%s %s", cond.Emoji, Bold(cond.Description)),
		fmt.Sprintf("%s (feels %s)",
			Yellow(units.FormatTemp(data.Current.Temperature, imperial)),
			units.FormatTemp(data.Current.ApparentTemperature, imperial)),
		fmt.Sprintf("Humidity: %s", Cyan(fmt.Sprintf("%d%%", data.Current.Humidity))),
		fmt.Sprintf("Wind: %s %s",
			Green(fmt.Sprintf("%.0f %s", data.Current.WindSpeed, units.WindUnit(imperial))),
			units.WindCardinal(data.Current.WindDirection)),
	}

	// Merge ASCII art lines and info lines side by side
	maxLines := len(artLines)
	if len(infoLines) > maxLines {
		maxLines = len(infoLines)
	}

	for i := 0; i < maxLines; i++ {
		artPart := ""
		if i < len(artLines) {
			artPart = artLines[i]
		}
		infoPart := ""
		if i < len(infoLines) {
			infoPart = infoLines[i]
		}
		// Pad art to fixed width
		for visLen(artPart) < 16 {
			artPart += " "
		}
		line := fmt.Sprintf("  %s%s", artPart, infoPart)
		b.WriteString(padLine(line))
	}

	b.WriteString(emptyLine())

	// Divider
	b.WriteString(divider())

	// Forecast header
	forecastHeader := fmt.Sprintf("  %-12s %5s %5s  %s", Dim("Day"), Dim("Hi"), Dim("Lo"), Dim("Cond."))
	b.WriteString(padLine(forecastHeader))

	// Daily forecast rows
	limit := days
	if limit > len(data.Daily) {
		limit = len(data.Daily)
	}

	for i := 0; i < limit; i++ {
		d := data.Daily[i]
		t, _ := time.Parse("2006-01-02", d.Date)
		dayName := t.Format("Mon 02")
		fc := GetCondition(d.WeatherCode)

		row := fmt.Sprintf("  %-12s %5s %5s  %s %s",
			dayName,
			units.FormatTemp(d.TemperatureMax, imperial),
			units.FormatTemp(d.TemperatureMin, imperial),
			fc.Emoji,
			fc.Description,
		)
		b.WriteString(padLine(row))
	}

	// Bottom border
	b.WriteString(bottomBorder())

	return b.String()
}

func topBorder() string {
	return fmt.Sprintf("\u250C%s\u2510\n", strings.Repeat("\u2500", cardWidth))
}

func bottomBorder() string {
	return fmt.Sprintf("\u2514%s\u2518\n", strings.Repeat("\u2500", cardWidth))
}

func divider() string {
	return fmt.Sprintf("\u251C%s\u2524\n", strings.Repeat("\u2500", cardWidth))
}

func emptyLine() string {
	return fmt.Sprintf("\u2502%s\u2502\n", strings.Repeat(" ", cardWidth))
}

func padLine(content string) string {
	// Calculate visible length (excluding ANSI codes)
	vl := visLen(content)
	padding := cardWidth - vl
	if padding < 0 {
		padding = 0
	}
	return fmt.Sprintf("\u2502%s%s\u2502\n", content, strings.Repeat(" ", padding))
}

// visLen returns the visible length of a string, excluding ANSI escape codes.
func visLen(s string) int {
	inEscape := false
	length := 0
	for _, r := range s {
		if r == '\033' {
			inEscape = true
			continue
		}
		if inEscape {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				inEscape = false
			}
			continue
		}
		length++
	}
	return length
}
