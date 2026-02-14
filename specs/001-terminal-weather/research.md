# Research: Terminal Weather Display

**Date**: 2026-02-14
**Feature**: 001-terminal-weather

## Decision 1: Weather API

**Decision**: Open-Meteo API (`api.open-meteo.com/v1/forecast`)

**Rationale**: Free, no API key required, supports both current weather and daily forecast in a single request. Provides server-side unit conversion (metric/imperial). Uses standard WMO weather codes.

**Alternatives considered**:
- OpenWeatherMap: Requires API key registration, rate-limited free tier
- wttr.in: Pre-formatted output, less control over display

**Key details**:
- Current weather endpoint: `https://api.open-meteo.com/v1/forecast?latitude={lat}&longitude={lon}&current=temperature_2m,relative_humidity_2m,apparent_temperature,wind_speed_10m,wind_direction_10m,weather_code`
- Daily forecast: same endpoint with `&daily=temperature_2m_max,temperature_2m_min,weather_code&timezone=auto&forecast_days={n}`
- Unit params: `temperature_unit=fahrenheit`, `wind_speed_unit=mph` for imperial
- Geocoding: `https://geocoding-api.open-meteo.com/v1/search?name={city}&count=1&language=en&format=json`

## Decision 2: Location Detection (Primary)

**Decision**: macOS CoreLocation via cgo with Objective-C bridge

**Rationale**: Native location access provides best accuracy. The async CLLocationManager callback pattern can be bridged to synchronous Go via CFRunLoopRun/CFRunLoopStop.

**Alternatives considered**:
- Swift helper binary: Two binaries to manage, more complex distribution
- IP-only: Less accurate, no native integration

**Key details**:
- cgo flags: `-x objective-c -mmacosx-version-min=10.14` (CFLAGS), `-framework CoreLocation -framework Foundation` (LDFLAGS)
- Authorization: macOS prompts automatically on `startUpdatingLocation()` (no explicit `requestWhenInUseAuthorization` needed)
- Timeout: Use `CFRunLoopRunInMode` with 10-second timeout to avoid hanging
- Info.plist: Embed via `-sectcreate __TEXT __info_plist ./Info.plist` linker flag with `NSLocationWhenInUseUsageDescription`
- Build: `CGO_LDFLAGS_ALLOW="-sectcreate"` environment variable required

## Decision 3: Location Detection (Fallback)

**Decision**: ip-api.com (`http://ip-api.com/json/`)

**Rationale**: Simplest free option, no API key required, returns lat/lon/city in a single call. 45 requests/minute rate limit is sufficient for a CLI tool.

**Alternatives considered**:
- ipapi.co: 1000/day limit, slightly more restrictive
- FreeIPAPI: Good rate limit but European-focused servers on free tier

**Key details**:
- Endpoint: `http://ip-api.com/json/` (returns location for caller's IP)
- Response fields: `lat`, `lon`, `city`, `regionName`, `country`
- Rate limit: 45 requests per minute

## Decision 4: WMO Weather Code Mapping

**Decision**: Static lookup table mapping WMO codes to descriptions, emoji, and ASCII art

**Rationale**: The code set is fixed and small (28 distinct codes). A static map is simpler and faster than any dynamic approach.

**Code categories**:
- 0-3: Clear to cloudy (sun, partly cloudy, overcast)
- 45-48: Fog
- 51-57: Drizzle (including freezing)
- 61-67: Rain (including freezing)
- 71-77: Snow
- 80-82: Rain showers
- 85-86: Snow showers
- 95-99: Thunderstorms

## Decision 5: Go Project Structure

**Decision**: Single-package Go module with internal packages for separation

**Rationale**: The tool is small enough that a flat structure with internal packages keeps it simple while maintaining testability. Using `internal/` prevents accidental external imports.

**Alternatives considered**:
- Single flat package: Too monolithic for testing
- Full cmd/pkg layout: Over-engineered for a CLI tool of this size
