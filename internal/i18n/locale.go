package i18n

import (
	"os"
	"os/exec"
	"strings"
)

// detectLocale returns the 2-letter language code from the system locale.
// Checks LC_ALL > LANG > LC_MESSAGES, then falls back to macOS AppleLanguages.
func detectLocale() string {
	for _, env := range []string{"LC_ALL", "LANG", "LC_MESSAGES"} {
		if val := os.Getenv(env); val != "" {
			if code := extractLangCode(val); code != "" {
				return code
			}
		}
	}
	return detectMacOSLocale()
}

// extractLangCode extracts a 2-letter language code from a locale string like "de_DE.UTF-8".
func extractLangCode(locale string) string {
	locale = strings.TrimSpace(locale)
	if locale == "" || locale == "C" || locale == "POSIX" {
		return ""
	}
	// Strip encoding (e.g., ".UTF-8")
	if idx := strings.Index(locale, "."); idx > 0 {
		locale = locale[:idx]
	}
	// Strip region (e.g., "_DE")
	if idx := strings.Index(locale, "_"); idx > 0 {
		locale = locale[:idx]
	}
	if len(locale) >= 2 {
		return strings.ToLower(locale[:2])
	}
	return ""
}

// detectMacOSLocale reads the macOS AppleLanguages preference.
func detectMacOSLocale() string {
	out, err := exec.Command("defaults", "read", "NSGlobalDomain", "AppleLanguages").Output()
	if err != nil {
		return "en"
	}
	// Output looks like:  (\n    "de-DE",\n    "en-US"\n)
	// Extract the first language code
	s := string(out)
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		line = strings.Trim(line, `",()`)
		line = strings.TrimSpace(line)
		if len(line) >= 2 {
			// Handle "de-DE" or "de"
			code := line
			if idx := strings.Index(code, "-"); idx > 0 {
				code = code[:idx]
			}
			return strings.ToLower(code[:2])
		}
	}
	return "en"
}
