package location

import "net/http"
import "time"

const ipGeoURL = "http://ip-api.com/json/"

// GetIPLocation fetches location from IP geolocation service.
func GetIPLocation() (Location, error) {
	return fetchIPLocation(ipGeoURL, &http.Client{Timeout: 5 * time.Second})
}
