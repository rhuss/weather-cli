package location

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchIPLocation(t *testing.T) {
	tests := []struct {
		name       string
		response   interface{}
		statusCode int
		wantErr    bool
		wantCity   string
		wantLat    float64
	}{
		{
			name: "successful response",
			response: map[string]interface{}{
				"status":     "success",
				"lat":        52.52,
				"lon":        13.41,
				"city":       "Berlin",
				"regionName": "Berlin",
				"country":    "Germany",
			},
			statusCode: 200,
			wantErr:    false,
			wantCity:   "Berlin",
			wantLat:    52.52,
		},
		{
			name: "failed status",
			response: map[string]interface{}{
				"status":  "fail",
				"message": "invalid query",
			},
			statusCode: 200,
			wantErr:    true,
		},
		{
			name:       "server error",
			response:   "error",
			statusCode: 500,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				json.NewEncoder(w).Encode(tt.response)
			}))
			defer server.Close()

			loc, err := fetchIPLocation(server.URL, server.Client())
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if loc.City != tt.wantCity {
				t.Errorf("city = %q, want %q", loc.City, tt.wantCity)
			}
			if loc.Latitude != tt.wantLat {
				t.Errorf("latitude = %f, want %f", loc.Latitude, tt.wantLat)
			}
			if loc.Source != "ip" {
				t.Errorf("source = %q, want %q", loc.Source, "ip")
			}
		})
	}
}

func TestResolveLocationManual(t *testing.T) {
	cfg := Config{Latitude: 48.85, Longitude: 2.35}
	loc, err := ResolveLocation(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loc.Source != "manual" {
		t.Errorf("source = %q, want %q", loc.Source, "manual")
	}
	if loc.Latitude != 48.85 {
		t.Errorf("latitude = %f, want 48.85", loc.Latitude)
	}
}

func TestResolveLocationCity(t *testing.T) {
	GeocodeFunc = func(city string) (float64, float64, string, string, error) {
		return 52.52, 13.41, "Berlin", "Germany", nil
	}
	defer func() { GeocodeFunc = nil }()

	cfg := Config{City: "Berlin"}
	loc, err := ResolveLocation(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loc.City != "Berlin" {
		t.Errorf("city = %q, want %q", loc.City, "Berlin")
	}
}

func TestResolveLocationLatLonOverCity(t *testing.T) {
	geocodeCalled := false
	GeocodeFunc = func(city string) (float64, float64, string, string, error) {
		geocodeCalled = true
		return 0, 0, "", "", nil
	}
	defer func() { GeocodeFunc = nil }()

	cfg := Config{City: "Berlin", Latitude: 48.85, Longitude: 2.35}
	loc, err := ResolveLocation(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if geocodeCalled {
		t.Error("geocode should not be called when lat/lon are provided")
	}
	if loc.Source != "manual" {
		t.Errorf("source = %q, want %q", loc.Source, "manual")
	}
}
