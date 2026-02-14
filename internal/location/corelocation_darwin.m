#import <Foundation/Foundation.h>
#import <CoreLocation/CoreLocation.h>
#include "corelocation_darwin.h"

@interface LocationDelegate : NSObject <CLLocationManagerDelegate>
@property (nonatomic, strong) CLLocation *lastLocation;
@property (nonatomic, strong) NSError *lastError;
@property (nonatomic) BOOL done;
@end

@implementation LocationDelegate

- (void)locationManager:(CLLocationManager *)manager
     didUpdateLocations:(NSArray<CLLocation *> *)locations {
    self.lastLocation = [locations lastObject];
    self.done = YES;
    CFRunLoopStop(CFRunLoopGetCurrent());
}

- (void)locationManager:(CLLocationManager *)manager
       didFailWithError:(NSError *)error {
    self.lastError = error;
    self.done = YES;
    CFRunLoopStop(CFRunLoopGetCurrent());
}

@end

int get_current_location(CLResult *result, double timeout_seconds) {
    @autoreleasepool {
        if (![CLLocationManager locationServicesEnabled]) {
            return -1;
        }

        CLLocationManager *manager = [[CLLocationManager alloc] init];
        LocationDelegate *delegate = [[LocationDelegate alloc] init];
        manager.delegate = delegate;
        manager.desiredAccuracy = kCLLocationAccuracyKilometer;

        [manager startUpdatingLocation];

        CFRunLoopRunInMode(kCFRunLoopDefaultMode, timeout_seconds, false);

        [manager stopUpdatingLocation];

        if (delegate.lastLocation != nil) {
            result->latitude = delegate.lastLocation.coordinate.latitude;
            result->longitude = delegate.lastLocation.coordinate.longitude;
            return 0;
        }

        if (delegate.lastError != nil) {
            return (int)delegate.lastError.code;
        }

        return -2; // timeout
    }
}

int location_services_enabled(void) {
    return [CLLocationManager locationServicesEnabled] ? 1 : 0;
}
