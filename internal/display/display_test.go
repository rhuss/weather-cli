package display

import (
	"goweather/internal/i18n"
	"goweather/internal/weather"
	"strings"
	"testing"
)

func init() {
	i18n.Init("en")
}

func TestRenderWeatherCard(t *testing.T) {
	data := &weather.WeatherData{
		Current: weather.CurrentWeather{
			Temperature:         18.5,
			ApparentTemperature: 16.2,
			Humidity:            55,
			WindSpeed:           12.0,
			WindDirection:       240,
			WeatherCode:        0,
			Time:               "2026-02-14T12:00",
		},
		Daily: []weather.DailyForecast{
			{Date: "2026-02-14", TemperatureMax: 20, TemperatureMin: 12, WeatherCode: 0},
			{Date: "2026-02-15", TemperatureMax: 18, TemperatureMin: 10, WeatherCode: 3},
			{Date: "2026-02-16", TemperatureMax: 15, TemperatureMin: 8, WeatherCode: 61},
		},
	}

	ColorEnabled = true
	output := RenderWeatherCard("Berlin, Germany", data, false, 3)

	// Should contain box-drawing characters
	if !strings.Contains(output, "\u250C") { // top-left corner
		t.Error("output missing top-left box corner")
	}
	if !strings.Contains(output, "\u2514") { // bottom-left corner
		t.Error("output missing bottom-left box corner")
	}

	// Should contain location
	if !strings.Contains(output, "Berlin") {
		t.Error("output missing location name")
	}

	// Should contain temperature
	if !strings.Contains(output, "18") {
		t.Error("output missing temperature value")
	}

	// Should contain weather emoji
	if !strings.Contains(output, "\u2600") { // sun
		t.Error("output missing sun emoji for clear sky")
	}
}

func TestRenderNoColor(t *testing.T) {
	data := &weather.WeatherData{
		Current: weather.CurrentWeather{
			Temperature:         18.5,
			ApparentTemperature: 16.2,
			Humidity:            55,
			WindSpeed:           12.0,
			WindDirection:       240,
			WeatherCode:        0,
		},
		Daily: []weather.DailyForecast{
			{Date: "2026-02-14", TemperatureMax: 20, TemperatureMin: 12, WeatherCode: 0},
		},
	}

	ColorEnabled = false
	output := RenderWeatherCard("Berlin", data, false, 1)

	// Should not contain ANSI escape codes
	if strings.Contains(output, "\033[") {
		t.Error("output contains ANSI escape codes in no-color mode")
	}

	// Should still contain content
	if !strings.Contains(output, "Berlin") {
		t.Error("output missing location in no-color mode")
	}

	// Reset
	ColorEnabled = true
}

func TestGetConditionUnknown(t *testing.T) {
	c := GetCondition(999)
	if c.Description != "Unknown" {
		t.Errorf("description = %q, want %q", c.Description, "Unknown")
	}
	if c.Emoji != "\u2753" { // question mark
		t.Errorf("emoji = %q, want question mark", c.Emoji)
	}
}

func TestVisLen(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"hello", 5},
		{"\033[1mhello\033[0m", 5},
		{"\033[31mred\033[0m text", 8},
		{"", 0},
	}

	for _, tt := range tests {
		got := visLen(tt.input)
		if got != tt.want {
			t.Errorf("visLen(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}
