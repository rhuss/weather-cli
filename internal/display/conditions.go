package display

// WeatherCondition maps a WMO code to display properties.
type WeatherCondition struct {
	Code        int
	Description string
	Emoji       string
	Category    string
}

var conditions = map[int]WeatherCondition{
	0:  {0, "Clear sky", "\u2600\uFE0F", "clear"},
	1:  {1, "Mainly clear", "\U0001F324\uFE0F", "clear"},
	2:  {2, "Partly cloudy", "\u26C5", "cloudy"},
	3:  {3, "Overcast", "\u2601\uFE0F", "cloudy"},
	45: {45, "Fog", "\U0001F32B\uFE0F", "fog"},
	48: {48, "Depositing rime fog", "\U0001F32B\uFE0F", "fog"},
	51: {51, "Light drizzle", "\U0001F326\uFE0F", "rain"},
	53: {53, "Moderate drizzle", "\U0001F326\uFE0F", "rain"},
	55: {55, "Dense drizzle", "\U0001F326\uFE0F", "rain"},
	56: {56, "Light freezing drizzle", "\U0001F327\uFE0F", "rain"},
	57: {57, "Dense freezing drizzle", "\U0001F327\uFE0F", "rain"},
	61: {61, "Slight rain", "\U0001F327\uFE0F", "rain"},
	63: {63, "Moderate rain", "\U0001F327\uFE0F", "rain"},
	65: {65, "Heavy rain", "\U0001F327\uFE0F", "rain"},
	66: {66, "Light freezing rain", "\U0001F327\uFE0F", "rain"},
	67: {67, "Heavy freezing rain", "\U0001F327\uFE0F", "rain"},
	71: {71, "Slight snow", "\u2744\uFE0F", "snow"},
	73: {73, "Moderate snow", "\u2744\uFE0F", "snow"},
	75: {75, "Heavy snow", "\u2744\uFE0F", "snow"},
	77: {77, "Snow grains", "\u2744\uFE0F", "snow"},
	80: {80, "Slight rain showers", "\U0001F326\uFE0F", "rain"},
	81: {81, "Moderate rain showers", "\U0001F326\uFE0F", "rain"},
	82: {82, "Violent rain showers", "\u26C8\uFE0F", "rain"},
	85: {85, "Slight snow showers", "\U0001F328\uFE0F", "snow"},
	86: {86, "Heavy snow showers", "\U0001F328\uFE0F", "snow"},
	95: {95, "Thunderstorm", "\u26A1", "storm"},
	96: {96, "Thunderstorm with slight hail", "\u26A1", "storm"},
	99: {99, "Thunderstorm with heavy hail", "\u26A1", "storm"},
}

// GetCondition returns the WeatherCondition for a WMO code.
// Returns an "Unknown" condition for unrecognized codes.
func GetCondition(code int) WeatherCondition {
	if c, ok := conditions[code]; ok {
		return c
	}
	return WeatherCondition{code, "Unknown", "\u2753", "unknown"}
}
