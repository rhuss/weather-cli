package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const geocodingURL = "https://geocoding-api.open-meteo.com/v1/search"

type geocodingResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Country   string  `json:"country"`
	} `json:"results"`
}

// GeocodeCity resolves a city name to coordinates using Open-Meteo geocoding.
func GeocodeCity(name string) (lat, lon float64, city, country string, err error) {
	return geocodeCityWithClient(name, &http.Client{Timeout: 5 * time.Second}, geocodingURL)
}

func geocodeCityWithClient(name string, client *http.Client, baseURL string) (lat, lon float64, city, country string, err error) {
	u := fmt.Sprintf("%s?name=%s&count=1&language=en&format=json", baseURL, url.QueryEscape(name))

	resp, err := client.Get(u)
	if err != nil {
		return 0, 0, "", "", fmt.Errorf("geocoding request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, "", "", fmt.Errorf("geocoding API returned status %d", resp.StatusCode)
	}

	var geoResp geocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return 0, 0, "", "", fmt.Errorf("failed to parse geocoding response: %w", err)
	}

	if len(geoResp.Results) == 0 {
		return 0, 0, "", "", fmt.Errorf("city not found: %s", name)
	}

	r := geoResp.Results[0]
	return r.Latitude, r.Longitude, r.Name, r.Country, nil
}
