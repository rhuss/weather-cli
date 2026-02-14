package units

import "fmt"

// TempUnit returns the temperature unit suffix.
func TempUnit(imperial bool) string {
	if imperial {
		return "°F"
	}
	return "°C"
}

// WindUnit returns the wind speed unit suffix.
func WindUnit(imperial bool) string {
	if imperial {
		return "mph"
	}
	return "km/h"
}

// FormatTemp formats a temperature value with its unit.
func FormatTemp(temp float64, imperial bool) string {
	return fmt.Sprintf("%.0f%s", temp, TempUnit(imperial))
}

// WindCardinal converts wind direction degrees to a cardinal direction.
func WindCardinal(degrees int) string {
	dirs := []string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE",
		"S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}
	idx := ((degrees + 11) / 22) % 16
	return dirs[idx]
}
