package weather

// CurrentWeather holds current weather conditions from the API.
type CurrentWeather struct {
	Temperature        float64
	ApparentTemperature float64
	Humidity           int
	WindSpeed          float64
	WindDirection      int
	WeatherCode        int
	Time               string
}

// DailyForecast holds one day's forecast data.
type DailyForecast struct {
	Date           string
	TemperatureMax float64
	TemperatureMin float64
	WeatherCode    int
}

// WeatherData bundles current conditions with the daily forecast.
type WeatherData struct {
	Current  CurrentWeather
	Daily    []DailyForecast
	Timezone string
}
