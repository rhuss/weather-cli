package location

/*
#cgo CFLAGS: -x objective-c -mmacosx-version-min=10.14
#cgo LDFLAGS: -framework CoreLocation -framework Foundation -mmacosx-version-min=10.14

#include "corelocation_darwin.h"
*/
import "C"
import "fmt"

const coreLocationTimeout = 10.0 // seconds

// GetCoreLocation attempts to get the current location via macOS CoreLocation.
func GetCoreLocation() (Location, error) {
	if C.location_services_enabled() == 0 {
		return Location{}, fmt.Errorf("location services disabled")
	}

	var result C.CLResult
	ret := C.get_current_location(&result, C.double(coreLocationTimeout))
	if ret != 0 {
		return Location{}, fmt.Errorf("CoreLocation failed with code %d", int(ret))
	}

	return Location{
		Latitude:  float64(result.latitude),
		Longitude: float64(result.longitude),
		Source:    "corelocation",
	}, nil
}
