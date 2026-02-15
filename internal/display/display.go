package display

import (
	"fmt"
	"goweather/internal/i18n"
	"goweather/internal/units"
	"goweather/internal/weather"
	"strings"
	"unicode"
	"unicode/utf8"
)

const cardWidth = 52

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
		fmt.Sprintf("%s (%s %s)",
			Yellow(units.FormatTemp(data.Current.Temperature, imperial)),
			i18n.Label("feels"),
			units.FormatTemp(data.Current.ApparentTemperature, imperial)),
		fmt.Sprintf("%s %s", i18n.Label("humidity"), Cyan(fmt.Sprintf("%d%%", data.Current.Humidity))),
		fmt.Sprintf("%s %s %s", i18n.Label("wind"),
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

	// Forecast table with fixed column positions
	// Columns: Day(10) Hi(6) Lo(6) Cond(rest)
	b.WriteString(padLine(forecastRow(Dim(i18n.Label("day")), Dim(i18n.Label("hi")), Dim(i18n.Label("lo")), Dim(i18n.Label("cond")), "")))

	// Daily forecast rows
	limit := days
	if limit > len(data.Daily) {
		limit = len(data.Daily)
	}

	for i := 0; i < limit; i++ {
		d := data.Daily[i]
		dayName := i18n.FormatDay(d.Date)
		fc := GetCondition(d.WeatherCode)

		row := forecastRow(
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

// forecastRow builds a forecast row with fixed column widths using visual padding.
func forecastRow(day, hi, lo, emoji, desc string) string {
	var b strings.Builder
	b.WriteString("  ")

	// Day column: 10 visible columns
	b.WriteString(day)
	for pad := 10 - visLen(day); pad > 0; pad-- {
		b.WriteByte(' ')
	}

	// Hi column: 6 visible columns, right-aligned
	hiStr := hi
	for pad := 6 - visLen(hiStr); pad > 0; pad-- {
		b.WriteByte(' ')
	}
	b.WriteString(hiStr)

	// Lo column: 6 visible columns, right-aligned
	loStr := lo
	for pad := 6 - visLen(loStr); pad > 0; pad-- {
		b.WriteByte(' ')
	}
	b.WriteString(loStr)

	// Condition column
	b.WriteString("  ")
	b.WriteString(emoji)
	if desc != "" {
		b.WriteByte(' ')
		b.WriteString(desc)
	}

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
	vl := visLen(content)
	if vl > cardWidth {
		content = truncateToWidth(content, cardWidth)
		vl = visLen(content)
	}
	padding := cardWidth - vl
	if padding < 0 {
		padding = 0
	}
	return fmt.Sprintf("\u2502%s%s\u2502\n", content, strings.Repeat(" ", padding))
}

// truncateToWidth truncates a string (preserving ANSI codes) to fit within maxWidth terminal columns.
func truncateToWidth(s string, maxWidth int) string {
	var result strings.Builder
	inEscape := false
	width := 0

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])

		if r == '\033' {
			inEscape = true
			result.WriteRune(r)
			i += size
			continue
		}
		if inEscape {
			result.WriteRune(r)
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				inEscape = false
			}
			i += size
			continue
		}

		w := runeWidth(r)
		if w == 0 {
			// Zero-width characters (variation selectors, combining marks)
			result.WriteRune(r)
			i += size
			continue
		}
		if width+w > maxWidth {
			break
		}
		width += w
		result.WriteRune(r)
		i += size
	}

	return result.String()
}

// visLen returns the visible terminal column width of a string, excluding ANSI escape codes.
func visLen(s string) int {
	inEscape := false
	width := 0
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
		width += runeWidth(r)
	}
	return width
}

// runeWidth returns the terminal column width of a rune.
// Most characters are 1 column, wide/fullwidth and emoji are 2, combining/zero-width are 0.
func runeWidth(r rune) int {
	// Variation selectors (zero-width)
	if r >= 0xFE00 && r <= 0xFE0F {
		return 0
	}
	// Combining marks (zero-width)
	if unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) {
		return 0
	}
	// Zero-width joiner/non-joiner
	if r == 0x200D || r == 0x200B || r == 0x200C || r == 0x200E || r == 0x200F {
		return 0
	}
	// Skin tone modifiers (zero-width, combine with previous emoji)
	if r >= 0x1F3FB && r <= 0x1F3FF {
		return 0
	}
	// Regional indicator symbols (flags): 2 wide each
	if r >= 0x1F1E6 && r <= 0x1F1FF {
		return 2
	}
	// Common emoji ranges: 2 wide
	if r >= 0x1F300 && r <= 0x1F9FF {
		return 2
	}
	if r >= 0x2600 && r <= 0x27BF {
		return 2
	}
	// Misc symbols that render wide
	if r == 0x26A1 || r == 0x26C5 || r == 0x26C8 || r == 0x2753 {
		return 2
	}
	// CJK and fullwidth characters
	if unicode.Is(unicode.Han, r) {
		return 2
	}
	if r >= 0xFF01 && r <= 0xFF60 {
		return 2
	}
	if r >= 0xFFE0 && r <= 0xFFE6 {
		return 2
	}
	return 1
}
