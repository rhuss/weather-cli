package display

import "fmt"

// ColorEnabled controls whether ANSI color codes are emitted.
var ColorEnabled = true

const (
	reset   = "\033[0m"
	bold    = "\033[1m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	cyan    = "\033[36m"
	white   = "\033[37m"
	dimmed  = "\033[2m"
)

// Bold wraps text in bold ANSI codes if color is enabled.
func Bold(s string) string {
	if !ColorEnabled {
		return s
	}
	return fmt.Sprintf("%s%s%s", bold, s, reset)
}

// Colored wraps text in the given ANSI color code if color is enabled.
func Colored(s string, color string) string {
	if !ColorEnabled {
		return s
	}
	return fmt.Sprintf("%s%s%s", color, s, reset)
}

// Yellow returns yellow-colored text.
func Yellow(s string) string { return Colored(s, yellow) }

// Blue returns blue-colored text.
func Blue(s string) string { return Colored(s, blue) }

// Cyan returns cyan-colored text.
func Cyan(s string) string { return Colored(s, cyan) }

// Red returns red-colored text.
func Red(s string) string { return Colored(s, red) }

// Green returns green-colored text.
func Green(s string) string { return Colored(s, green) }

// Dim returns dimmed text.
func Dim(s string) string { return Colored(s, dimmed) }
