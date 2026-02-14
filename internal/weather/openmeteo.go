package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://api.open-meteo.com/v1/forecast"

// Client fetches weather data from the Open-Meteo API.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// NewClient creates a weather API client with default settings.
func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		BaseURL:    baseURL,
	}
}

// apiResponse mirrors the Open-Meteo JSON structure.
type apiResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Current   struct {
		Time               string  `json:"time"`
		Temperature2m      float64 `json:"temperature_2m"`
		RelativeHumidity2m int     `json:"relative_humidity_2m"`
		ApparentTemp       float64 `json:"apparent_temperature"`
		WindSpeed10m       float64 `json:"wind_speed_10m"`
		WindDirection10m   int     `json:"wind_direction_10m"`
		WeatherCode        int     `json:"weather_code"`
	} `json:"current"`
	Daily struct {
		Time           []string  `json:"time"`
		TempMax        []float64 `json:"temperature_2m_max"`
		TempMin        []float64 `json:"temperature_2m_min"`
		WeatherCode    []int     `json:"weather_code"`
	} `json:"daily"`
}

// FetchWeather retrieves current weather and daily forecast.
func (c *Client) FetchWeather(lat, lon float64, days int, imperial bool) (*WeatherData, error) {
	tempUnit := "celsius"
	windUnit := "kmh"
	if imperial {
		tempUnit = "fahrenheit"
		windUnit = "mph"
	}

	url := fmt.Sprintf(
		"%s?latitude=%.4f&longitude=%.4f"+
			"&current=temperature_2m,relative_humidity_2m,apparent_temperature,wind_speed_10m,wind_direction_10m,weather_code"+
			"&daily=temperature_2m_max,temperature_2m_min,weather_code"+
			"&timezone=auto&forecast_days=%d"+
			"&temperature_unit=%s&wind_speed_unit=%s",
		c.BaseURL, lat, lon, days, tempUnit, windUnit,
	)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("weather API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API returned status %d", resp.StatusCode)
	}

	var apiResp apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse weather response: %w", err)
	}

	current := CurrentWeather{
		Temperature:         apiResp.Current.Temperature2m,
		ApparentTemperature: apiResp.Current.ApparentTemp,
		Humidity:            apiResp.Current.RelativeHumidity2m,
		WindSpeed:           apiResp.Current.WindSpeed10m,
		WindDirection:       apiResp.Current.WindDirection10m,
		WeatherCode:         apiResp.Current.WeatherCode,
		Time:                apiResp.Current.Time,
	}

	daily := make([]DailyForecast, len(apiResp.Daily.Time))
	for i := range apiResp.Daily.Time {
		daily[i] = DailyForecast{
			Date:           apiResp.Daily.Time[i],
			TemperatureMax: apiResp.Daily.TempMax[i],
			TemperatureMin: apiResp.Daily.TempMin[i],
			WeatherCode:    apiResp.Daily.WeatherCode[i],
		}
	}

	return &WeatherData{
		Current:  current,
		Daily:    daily,
		Timezone: apiResp.Timezone,
	}, nil
}
