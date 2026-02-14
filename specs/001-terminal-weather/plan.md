# Implementation Plan: Terminal Weather Display

**Branch**: `001-terminal-weather` | **Date**: 2026-02-14 | **Spec**: [spec.md](spec.md)
**Input**: Feature specification from `specs/001-terminal-weather/spec.md`

## Summary

A Go CLI tool (`goweather`) that displays current weather conditions and a multi-day forecast in the terminal with ASCII art and emoji. Uses macOS CoreLocation (via cgo) for automatic location detection with IP geolocation fallback, and Open-Meteo API for weather data. Outputs a compact bordered card with Unicode box-drawing characters.

## Technical Context

**Language/Version**: Go 1.22+ (cgo enabled)
**Primary Dependencies**: CoreLocation framework (cgo), Open-Meteo API, ip-api.com
**Storage**: N/A (no persistent storage)
**Testing**: `go test` with table-driven tests, mock HTTP responses
**Target Platform**: macOS 10.14+ (Apple Silicon and Intel)
**Project Type**: Single CLI binary
**Performance Goals**: Weather displayed within 3 seconds of invocation
**Constraints**: Single binary, no external runtime dependencies, minimum 40-column terminal
**Scale/Scope**: Single-user CLI tool, one request per invocation

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

Constitution is a template (not yet configured for this project). No gates to check. PASS.

## Project Structure

### Documentation (this feature)

```text
specs/001-terminal-weather/
├── spec.md              # Feature specification
├── plan.md              # This file
├── research.md          # Phase 0: API research, technology decisions
├── data-model.md        # Phase 1: Entity definitions
├── quickstart.md        # Phase 1: Build and run instructions
└── tasks.md             # Phase 2: Task breakdown (via /speckit.tasks)
```

### Source Code (repository root)

```text
.
├── main.go                      # Entry point, flag parsing, orchestration
├── go.mod                       # Go module definition
├── go.sum                       # Dependency checksums
├── Info.plist                   # macOS location permission description
├── Makefile                     # Build commands (build, clean, test)
├── internal/
│   ├── location/
│   │   ├── location.go          # Location interface and types
│   │   ├── corelocation_darwin.go   # CoreLocation cgo bridge (Go side)
│   │   ├── corelocation_darwin.h    # CoreLocation C header
│   │   ├── corelocation_darwin.m    # CoreLocation Objective-C implementation
│   │   ├── ipgeo.go             # IP geolocation fallback
│   │   └── location_test.go     # Location tests (mocked)
│   ├── weather/
│   │   ├── weather.go           # Weather types and interface
│   │   ├── openmeteo.go         # Open-Meteo API client
│   │   ├── geocoding.go         # Open-Meteo geocoding (city to lat/lon)
│   │   └── weather_test.go      # Weather API tests (mocked HTTP)
│   ├── display/
│   │   ├── display.go           # Card rendering, box drawing, layout
│   │   ├── ascii.go             # ASCII art weather icons
│   │   ├── conditions.go        # WMO code to description/emoji mapping
│   │   ├── color.go             # ANSI color helpers with --no-color support
│   │   └── display_test.go      # Display rendering tests
│   └── units/
│       ├── units.go             # Unit conversion (metric/imperial)
│       └── units_test.go        # Unit conversion tests
└── testdata/
    ├── weather_response.json    # Sample Open-Meteo response for tests
    └── geocoding_response.json  # Sample geocoding response for tests
```

**Structure Decision**: Single Go module with `internal/` packages for logical separation. The `internal/` prefix prevents external imports while keeping packages small and testable. Four packages: `location` (detection), `weather` (API client), `display` (rendering), `units` (conversion).

## Key Architecture Decisions

### CoreLocation Bridge Pattern

The Objective-C bridge uses CFRunLoopRunInMode with a 10-second timeout to convert CoreLocation's async delegate callbacks into a synchronous C function call. Go calls this C function via cgo. The bridge consists of three files:
- `.h` header defining the C interface (Location struct, get_current_location function)
- `.m` implementation with CLLocationManagerDelegate
- `.go` file with cgo directives and Go wrapper functions

### Info.plist Embedding

The `NSLocationWhenInUseUsageDescription` plist is embedded into the binary at link time using `-sectcreate __TEXT __info_plist ./Info.plist`. This requires `CGO_LDFLAGS_ALLOW="-sectcreate"` at build time.

### Weather Data Flow

```
main.go
  ├─ Parse flags (--city, --lat/--lon, --imperial, --no-color, --days)
  ├─ Resolve location:
  │   ├─ If --lat/--lon provided → use directly
  │   ├─ If --city provided → geocode via Open-Meteo
  │   └─ Else → CoreLocation → fallback to IP geolocation
  ├─ Fetch weather from Open-Meteo (current + daily)
  └─ Render display card to stdout
```

### Unit Handling

Units are requested server-side from Open-Meteo (`temperature_unit`, `wind_speed_unit` parameters) rather than converting client-side. This simplifies the code and avoids rounding errors.

## Complexity Tracking

No constitution violations to justify.
