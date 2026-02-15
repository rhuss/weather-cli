package i18n

import (
	"fmt"
	"strings"
	"time"
)

var active *Lang

// Init sets the active language. If langOverride is empty, locale detection is used.
// Falls back to English if the requested language is not available.
func Init(langOverride string) {
	code := langOverride
	if code == "" {
		code = detectLocale()
	}
	code = strings.ToLower(code)
	if l, ok := registry[code]; ok {
		active = l
		return
	}
	// Fallback to English
	active = registry["en"]
}

// Label returns a translated display label by key.
func Label(key string) string {
	if active == nil {
		return key
	}
	switch key {
	case "day":
		return active.LabelDay
	case "hi":
		return active.LabelHi
	case "lo":
		return active.LabelLo
	case "cond":
		return active.LabelCond
	case "humidity":
		return active.LabelHumidity
	case "wind":
		return active.LabelWind
	case "feels":
		return active.LabelFeels
	default:
		return key
	}
}

// Condition returns the translated weather condition description for a WMO code.
// Returns an empty string if no translation is available.
func Condition(code int) string {
	if active == nil {
		return ""
	}
	if desc, ok := active.Conditions[code]; ok {
		return desc
	}
	return ""
}

// DayAbbr returns the translated day abbreviation for a time.Weekday.
func DayAbbr(wd time.Weekday) string {
	if active == nil {
		return wd.String()[:3]
	}
	return active.DayAbbreviations[wd]
}

// Cardinal returns the translated cardinal direction for index 0-15.
func Cardinal(idx int) string {
	if active == nil || idx < 0 || idx > 15 {
		return "?"
	}
	return active.Cardinals[idx]
}

// FormatDay formats a date string (YYYY-MM-DD) as a localized "DayAbbr DD" string.
func FormatDay(dateStr string) string {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return dateStr
	}
	return fmt.Sprintf("%s %02d", DayAbbr(t.Weekday()), t.Day())
}

// TipManualLocation returns the localized tip message for manual location.
func TipManualLocation() string {
	if active == nil {
		return "Tip: Use --city or --lat/--lon to specify a location manually"
	}
	return active.TipManualLocation
}
