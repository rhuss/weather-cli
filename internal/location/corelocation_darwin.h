#ifndef CORELOCATION_DARWIN_H
#define CORELOCATION_DARWIN_H

typedef struct {
    double latitude;
    double longitude;
} CLResult;

// get_current_location attempts to get the current location via CoreLocation.
// Returns 0 on success, non-zero on failure.
// timeout_seconds specifies how long to wait for a location fix.
int get_current_location(CLResult *result, double timeout_seconds);

// location_services_enabled returns 1 if location services are enabled, 0 otherwise.
int location_services_enabled(void);

#endif
