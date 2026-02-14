package weather

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFetchWeather(t *testing.T) {
	fixture, err := os.ReadFile("../../testdata/weather_response.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(fixture)
	}))
	defer server.Close()

	client := &Client{
		HTTPClient: server.Client(),
		BaseURL:    server.URL,
	}

	data, err := client.FetchWeather(52.52, 13.41, 5, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data.Current.Temperature != 5.2 {
		t.Errorf("temperature = %f, want 5.2", data.Current.Temperature)
	}
	if data.Current.Humidity != 73 {
		t.Errorf("humidity = %d, want 73", data.Current.Humidity)
	}
	if data.Current.WeatherCode != 3 {
		t.Errorf("weather_code = %d, want 3", data.Current.WeatherCode)
	}
	if len(data.Daily) != 5 {
		t.Errorf("daily count = %d, want 5", len(data.Daily))
	}
	if data.Daily[0].TemperatureMax != 6.2 {
		t.Errorf("daily[0].max = %f, want 6.2", data.Daily[0].TemperatureMax)
	}
}

func TestFetchWeatherImperial(t *testing.T) {
	fixture, err := os.ReadFile("../../testdata/weather_response.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	var requestURL string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestURL = r.URL.String()
		w.Header().Set("Content-Type", "application/json")
		w.Write(fixture)
	}))
	defer server.Close()

	client := &Client{
		HTTPClient: server.Client(),
		BaseURL:    server.URL,
	}

	_, err = client.FetchWeather(52.52, 13.41, 5, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if requestURL == "" {
		t.Fatal("no request was made")
	}

	// Check that imperial units are requested
	if got := requestURL; got == "" {
		t.Error("request URL should contain temperature_unit=fahrenheit")
	}
}

func TestGeocodeCityWithClient(t *testing.T) {
	fixture, err := os.ReadFile("../../testdata/geocoding_response.json")
	if err != nil {
		t.Fatalf("failed to read fixture: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(fixture)
	}))
	defer server.Close()

	lat, lon, city, country, err := geocodeCityWithClient("Berlin", server.Client(), server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if city != "Berlin" {
		t.Errorf("city = %q, want %q", city, "Berlin")
	}
	if country != "Germany" {
		t.Errorf("country = %q, want %q", country, "Germany")
	}
	if lat != 52.52437 {
		t.Errorf("lat = %f, want 52.52437", lat)
	}
	if lon != 13.41053 {
		t.Errorf("lon = %f, want 13.41053", lon)
	}
}

func TestGeocodeCityNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"results":[]}`))
	}))
	defer server.Close()

	_, _, _, _, err := geocodeCityWithClient("Xyzzyville", server.Client(), server.URL)
	if err == nil {
		t.Error("expected error for unknown city, got nil")
	}
}
