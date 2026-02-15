package i18n

import (
	"os"
	"testing"
	"time"
)

func TestInitEnglish(t *testing.T) {
	Init("en")
	if active == nil {
		t.Fatal("active language is nil after Init(\"en\")")
	}
	if active.Code != "en" {
		t.Errorf("active.Code = %q, want %q", active.Code, "en")
	}
}

func TestInitGerman(t *testing.T) {
	Init("de")
	if active.Code != "de" {
		t.Errorf("active.Code = %q, want %q", active.Code, "de")
	}
}

func TestInitFallback(t *testing.T) {
	Init("xx") // unknown language
	if active.Code != "en" {
		t.Errorf("expected fallback to en, got %q", active.Code)
	}
}

func TestInitCaseInsensitive(t *testing.T) {
	Init("DE")
	if active.Code != "de" {
		t.Errorf("active.Code = %q, want %q", active.Code, "de")
	}
}

func TestLabel(t *testing.T) {
	Init("en")
	tests := []struct {
		key  string
		want string
	}{
		{"day", "Day"},
		{"hi", "Hi"},
		{"lo", "Lo"},
		{"cond", "Cond."},
		{"humidity", "Humidity:"},
		{"wind", "Wind:"},
		{"feels", "feels"},
	}
	for _, tt := range tests {
		got := Label(tt.key)
		if got != tt.want {
			t.Errorf("Label(%q) = %q, want %q", tt.key, got, tt.want)
		}
	}
}

func TestLabelGerman(t *testing.T) {
	Init("de")
	if got := Label("feels"); got != "gefühlt" {
		t.Errorf("Label(\"feels\") = %q, want %q", got, "gefühlt")
	}
	if got := Label("humidity"); got != "Feuchte:" {
		t.Errorf("Label(\"humidity\") = %q, want %q", got, "Feuchte:")
	}
}

func TestCondition(t *testing.T) {
	Init("en")
	if got := Condition(0); got != "Clear sky" {
		t.Errorf("Condition(0) = %q, want %q", got, "Clear sky")
	}
	if got := Condition(65); got != "Heavy rain" {
		t.Errorf("Condition(65) = %q, want %q", got, "Heavy rain")
	}
	// Unknown code
	if got := Condition(9999); got != "" {
		t.Errorf("Condition(9999) = %q, want empty string", got)
	}
}

func TestConditionGerman(t *testing.T) {
	Init("de")
	if got := Condition(0); got != "Klarer Himmel" {
		t.Errorf("Condition(0) = %q, want %q", got, "Klarer Himmel")
	}
}

func TestDayAbbr(t *testing.T) {
	Init("en")
	if got := DayAbbr(time.Monday); got != "Mon" {
		t.Errorf("DayAbbr(Monday) = %q, want %q", got, "Mon")
	}
	if got := DayAbbr(time.Sunday); got != "Sun" {
		t.Errorf("DayAbbr(Sunday) = %q, want %q", got, "Sun")
	}
}

func TestDayAbbrGerman(t *testing.T) {
	Init("de")
	if got := DayAbbr(time.Monday); got != "Mo" {
		t.Errorf("DayAbbr(Monday) = %q, want %q", got, "Mo")
	}
}

func TestCardinal(t *testing.T) {
	Init("en")
	if got := Cardinal(0); got != "N" {
		t.Errorf("Cardinal(0) = %q, want %q", got, "N")
	}
	if got := Cardinal(4); got != "E" {
		t.Errorf("Cardinal(4) = %q, want %q", got, "E")
	}
}

func TestCardinalGerman(t *testing.T) {
	Init("de")
	if got := Cardinal(4); got != "O" {
		t.Errorf("Cardinal(4) = %q, want %q", got, "O")
	}
}

func TestFormatDay(t *testing.T) {
	Init("en")
	// 2026-02-14 is a Saturday
	got := FormatDay("2026-02-14")
	if got != "Sat 14" {
		t.Errorf("FormatDay(\"2026-02-14\") = %q, want %q", got, "Sat 14")
	}
}

func TestFormatDayGerman(t *testing.T) {
	Init("de")
	got := FormatDay("2026-02-14")
	if got != "Sa 14" {
		t.Errorf("FormatDay(\"2026-02-14\") = %q, want %q", got, "Sa 14")
	}
}

func TestTipManualLocation(t *testing.T) {
	Init("en")
	got := TipManualLocation()
	want := "Tip: Use --city or --lat/--lon to specify a location manually"
	if got != want {
		t.Errorf("TipManualLocation() = %q, want %q", got, want)
	}
}

func TestLocaleDetectionFromEnv(t *testing.T) {
	orig := os.Getenv("LANG")
	defer os.Setenv("LANG", orig)

	os.Setenv("LANG", "de_DE.UTF-8")
	Init("")
	if active.Code != "de" {
		t.Errorf("expected auto-detect de from LANG, got %q", active.Code)
	}
}

func TestExtractLangCode(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"de_DE.UTF-8", "de"},
		{"en_US.UTF-8", "en"},
		{"fr", "fr"},
		{"zh_CN", "zh"},
		{"C", ""},
		{"POSIX", ""},
		{"", ""},
	}
	for _, tt := range tests {
		got := extractLangCode(tt.input)
		if got != tt.want {
			t.Errorf("extractLangCode(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestAllLanguagesRegistered(t *testing.T) {
	for _, code := range []string{"en", "de", "es", "fr", "it", "zh"} {
		if _, ok := registry[code]; !ok {
			t.Errorf("language %q not registered", code)
		}
	}
}

func TestAllLanguagesHaveConditions(t *testing.T) {
	codes := []int{0, 1, 2, 3, 45, 48, 51, 53, 55, 56, 57, 61, 63, 65, 66, 67, 71, 73, 75, 77, 80, 81, 82, 85, 86, 95, 96, 99}
	for langCode, lang := range registry {
		for _, wmo := range codes {
			if _, ok := lang.Conditions[wmo]; !ok {
				t.Errorf("language %q missing condition for WMO code %d", langCode, wmo)
			}
		}
	}
}
