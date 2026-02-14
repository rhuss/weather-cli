package units

import "testing"

func TestTempUnit(t *testing.T) {
	if got := TempUnit(false); got != "°C" {
		t.Errorf("TempUnit(false) = %q, want %q", got, "°C")
	}
	if got := TempUnit(true); got != "°F" {
		t.Errorf("TempUnit(true) = %q, want %q", got, "°F")
	}
}

func TestWindUnit(t *testing.T) {
	if got := WindUnit(false); got != "km/h" {
		t.Errorf("WindUnit(false) = %q, want %q", got, "km/h")
	}
	if got := WindUnit(true); got != "mph" {
		t.Errorf("WindUnit(true) = %q, want %q", got, "mph")
	}
}

func TestFormatTemp(t *testing.T) {
	tests := []struct {
		temp     float64
		imperial bool
		want     string
	}{
		{18.5, false, "18°C"},
		{65.3, true, "65°F"},
		{0, false, "0°C"},
		{-5.7, false, "-6°C"},
	}

	for _, tt := range tests {
		got := FormatTemp(tt.temp, tt.imperial)
		if got != tt.want {
			t.Errorf("FormatTemp(%f, %v) = %q, want %q", tt.temp, tt.imperial, got, tt.want)
		}
	}
}

func TestWindCardinal(t *testing.T) {
	tests := []struct {
		degrees int
		want    string
	}{
		{0, "N"},
		{90, "E"},
		{180, "S"},
		{270, "W"},
		{45, "NE"},
		{135, "SE"},
		{225, "SW"},
		{315, "NW"},
	}

	for _, tt := range tests {
		got := WindCardinal(tt.degrees)
		if got != tt.want {
			t.Errorf("WindCardinal(%d) = %q, want %q", tt.degrees, got, tt.want)
		}
	}
}
