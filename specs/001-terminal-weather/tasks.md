# Tasks: Terminal Weather Display

**Input**: Design documents from `/specs/001-terminal-weather/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md

**Tests**: Included as part of implementation tasks (table-driven Go tests alongside source files).

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and Go module structure

- [x] T001 (sdd-smoke-test-b6w.1) Initialize Go module with `go mod init` and create project directory structure per plan.md (`internal/location/`, `internal/weather/`, `internal/display/`, `internal/units/`, `testdata/`)
- [x] T002 (sdd-smoke-test-b6w.2) Create Info.plist with `NSLocationWhenInUseUsageDescription` for macOS location permission in `Info.plist`
- [x] T003 (sdd-smoke-test-b6w.3) Create Makefile with `build`, `test`, and `clean` targets (build target must use `CGO_LDFLAGS_ALLOW` and embed Info.plist) in `Makefile`
- [x] T004 (sdd-smoke-test-b6w.4) Create .gitignore with Go patterns (`*.exe`, `*.test`, `vendor/`, `*.out`, `goweather` binary) in `.gitignore`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core types and shared infrastructure that ALL user stories depend on

**CRITICAL**: No user story work can begin until this phase is complete

- [x] T005 (sdd-smoke-test-1ai.1) Define Location struct (Latitude, Longitude, City, Country, Source) and Config struct (City, Latitude, Longitude, Imperial, NoColor, Days) in `internal/location/location.go`
- [x] T006 (sdd-smoke-test-1ai.2) [P] Define CurrentWeather and DailyForecast structs matching Open-Meteo API response shape in `internal/weather/weather.go`
- [x] T007 (sdd-smoke-test-1ai.3) [P] Implement WMO weather code to WeatherCondition mapping (all 28 codes from data-model.md) with description, emoji, ASCII art, and category in `internal/display/conditions.go`
- [x] T008 (sdd-smoke-test-1ai.4) [P] Implement ANSI color helper functions with `--no-color` support (color on/off toggle, color wrapper functions) in `internal/display/color.go`
- [x] T009 (sdd-smoke-test-1ai.5) [P] Create sample Open-Meteo API response JSON fixture for tests in `testdata/weather_response.json`
- [x] T010 (sdd-smoke-test-1ai.6) [P] Create sample geocoding API response JSON fixture for tests in `testdata/geocoding_response.json`

**Checkpoint**: Foundation ready. All types defined, weather code mapping complete, test fixtures in place.

---

## Phase 3: User Story 1 - Quick Weather Check (Priority: P1) MVP

**Goal**: User runs `goweather` with no arguments, gets current weather and 5-day forecast for auto-detected location in a compact card.

**Independent Test**: Run the built binary with no args, verify a bordered weather card with current conditions and 5-day forecast is displayed.

### Implementation for User Story 1

- [x] T011 (sdd-smoke-test-so9.1) [P] [US1] Implement CoreLocation Objective-C bridge: C header defining Location struct and `get_current_location` function in `internal/location/corelocation_darwin.h`
- [x] T012 (sdd-smoke-test-so9.2) [P] [US1] Implement CoreLocation Objective-C bridge: CLLocationManagerDelegate with CFRunLoopRunInMode timeout (10s), error handling for denied/unavailable in `internal/location/corelocation_darwin.m`
- [x] T013 (sdd-smoke-test-so9.3) [US1] Implement Go wrapper for CoreLocation cgo bridge with CFLAGS/LDFLAGS directives, `GetCoreLocation()` function returning Location in `internal/location/corelocation_darwin.go`
- [x] T014 (sdd-smoke-test-so9.4) [US1] Implement IP geolocation fallback using ip-api.com (`http://ip-api.com/json/`), parse lat/lon/city/country from JSON response in `internal/location/ipgeo.go`
- [x] T015 (sdd-smoke-test-so9.5) [US1] Implement `ResolveLocation(cfg Config)` function: try CoreLocation first, fall back to IP geolocation silently on failure in `internal/location/location.go`
- [x] T016 (sdd-smoke-test-so9.6) [US1] Implement Open-Meteo API client: `FetchWeather(lat, lon float64, days int, imperial bool)` returning CurrentWeather and []DailyForecast, with unit params for metric/imperial in `internal/weather/openmeteo.go`
- [x] T017 (sdd-smoke-test-so9.7) [P] [US1] Implement ASCII art weather icons (sun, clouds, rain, snow, thunderstorm, fog, drizzle, unknown) as multi-line string constants in `internal/display/ascii.go`
- [x] T018 (sdd-smoke-test-so9.8) [US1] Implement card renderer: bordered box with location header, ASCII art + current conditions (temp, feels-like, humidity, wind), and daily forecast table using Unicode box-drawing in `internal/display/display.go`
- [x] T019 (sdd-smoke-test-so9.9) [US1] Implement main.go: parse flags (--city, --lat, --lon, --imperial, --no-color, --days), orchestrate location resolution, weather fetch, and display rendering in `main.go`
- [x] T020 (sdd-smoke-test-so9.10) [US1] Write tests for IP geolocation client using httptest mock server in `internal/location/location_test.go`
- [x] T021 (sdd-smoke-test-so9.11) [US1] Write tests for Open-Meteo client using httptest mock server and testdata fixtures in `internal/weather/weather_test.go`
- [x] T022 (sdd-smoke-test-so9.12) [US1] Write tests for display rendering (verify box-drawing output, verify no-color mode strips ANSI codes) in `internal/display/display_test.go`

**Checkpoint**: `goweather` with no arguments displays weather for auto-detected location. MVP is functional.

---

## Phase 4: User Story 2 - Manual Location Override (Priority: P2)

**Goal**: User can specify `--city Berlin` or `--lat 52.52 --lon 13.41` to get weather for a specific location.

**Independent Test**: Run `goweather --city Berlin` and verify Berlin weather is displayed. Run with invalid city and verify error message.

### Implementation for User Story 2

- [x] T023 (sdd-smoke-test-kz4.1) [US2] Implement Open-Meteo geocoding client: `GeocodeCity(name string)` returning lat/lon/city/country from geocoding API in `internal/weather/geocoding.go`
- [x] T024 (sdd-smoke-test-kz4.2) [US2] Update `ResolveLocation(cfg Config)` to handle --city (geocode), --lat/--lon (direct use), validate both lat+lon provided together, prioritize lat/lon over city in `internal/location/location.go`
- [x] T025 (sdd-smoke-test-kz4.3) [US2] Add error handling for invalid city name ("City not found") and missing lat/lon pair ("Both --lat and --lon must be provided together") in `main.go`
- [x] T026 (sdd-smoke-test-kz4.4) [US2] Write tests for geocoding client using httptest mock server and testdata fixture in `internal/weather/weather_test.go`
- [x] T027 (sdd-smoke-test-kz4.5) [US2] Write tests for location resolution priority logic (lat/lon > city > auto) in `internal/location/location_test.go`

**Checkpoint**: `goweather --city Berlin` and `goweather --lat 52.52 --lon 13.41` both work correctly.

---

## Phase 5: User Story 3 - Unit and Display Preferences (Priority: P3)

**Goal**: User can switch between metric and imperial units, and disable colors for piping.

**Independent Test**: Run `goweather --imperial` and verify Fahrenheit/mph output. Run `goweather --no-color | grep -c $'\033'` and verify zero escape codes.

### Implementation for User Story 3

- [x] T028 (sdd-smoke-test-5jk.1) [US3] Implement unit label helpers: return correct unit suffix (C/F, km/h/mph) and wind direction cardinal (N/S/E/W) based on imperial flag in `internal/units/units.go`
- [x] T029 (sdd-smoke-test-5jk.2) [US3] Update display renderer to use unit labels from units package, pass NoColor flag to color helpers in `internal/display/display.go`
- [x] T030 (sdd-smoke-test-5jk.3) [US3] Write tests for unit label functions (metric vs imperial outputs) in `internal/units/units_test.go`

**Checkpoint**: `goweather --imperial` shows Fahrenheit, `goweather --no-color` produces clean output.

---

## Phase 6: User Story 4 - Custom Forecast Range (Priority: P3)

**Goal**: User can specify `--days 3` to control forecast length.

**Independent Test**: Run `goweather --days 3` and verify only 3 forecast days are shown.

### Implementation for User Story 4

- [x] T031 (sdd-smoke-test-5ut.1) [US4] Add --days flag validation in main.go: enforce range 1-7, display error for out-of-range values in `main.go`
- [x] T032 (sdd-smoke-test-5ut.2) [US4] Update display renderer to respect days count when rendering forecast table in `internal/display/display.go`

**Checkpoint**: `goweather --days 3` shows exactly 3 forecast days, `goweather --days 10` shows error.

---

## Phase 6b: User Story 5 - Multi-Language Support (Priority: P3)

**Goal**: User can specify `--lang de` or rely on system locale for translated output.

**Independent Test**: Run `goweather --lang de` and verify German labels, day names, and condition descriptions.

### Implementation for User Story 5

- [x] T037 [US5] Define Lang struct with translatable fields (labels, day abbreviations, cardinals, conditions) in `internal/i18n/lang.go`
- [x] T038 [P] [US5] Implement English language pack in `internal/i18n/lang_en.go`
- [x] T039 [P] [US5] Implement German language pack in `internal/i18n/lang_de.go`
- [x] T040 [P] [US5] Implement Spanish language pack in `internal/i18n/lang_es.go`
- [x] T041 [P] [US5] Implement French language pack in `internal/i18n/lang_fr.go`
- [x] T042 [P] [US5] Implement Italian language pack in `internal/i18n/lang_it.go`
- [x] T043 [P] [US5] Implement Chinese language pack in `internal/i18n/lang_zh.go`
- [x] T044 [US5] Implement i18n Init, Label, Condition, DayAbbr, Cardinal, FormatDay functions with locale detection fallback in `internal/i18n/i18n.go`
- [x] T045 [US5] Implement system locale detection in `internal/i18n/locale.go`
- [x] T046 [US5] Add `--lang` flag to main.go and wire i18n.Init before output in `main.go`
- [x] T047 [US5] Integrate i18n labels into display renderer (labels, day formatting, cardinal directions) in `internal/display/display.go`
- [x] T048 [US5] Write tests for i18n language selection and fallback in `internal/i18n/i18n_test.go`

**Checkpoint**: `goweather --lang de` shows German output, system locale auto-detection works, unknown languages fall back to English.

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Edge cases, error handling, and build verification

- [x] T033 (sdd-smoke-test-14l.1) Handle edge case: unknown WMO weather codes display "Unknown" with question mark emoji in `internal/display/conditions.go`
- [x] T034 (sdd-smoke-test-14l.2) Handle edge case: network errors produce clear user-facing messages ("No network connection", "Unable to fetch weather data") in `main.go`
- [x] T035 (sdd-smoke-test-14l.3) [P] Verify build produces single binary with embedded Info.plist using `make build` and `otool -s __TEXT __info_plist`
- [x] T036 (sdd-smoke-test-14l.4) Run full test suite with `go test ./...` and verify all tests pass

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies, start immediately
- **Foundational (Phase 2)**: Depends on Setup completion, BLOCKS all user stories
- **US1 (Phase 3)**: Depends on Foundational. Core MVP, should complete first
- **US2 (Phase 4)**: Depends on Foundational + US1 (extends location resolution)
- **US3 (Phase 5)**: Depends on Foundational + US1 (extends display rendering)
- **US4 (Phase 6)**: Depends on Foundational + US1 (extends display rendering)
- **Polish (Phase 7)**: Depends on all user stories being complete

### User Story Dependencies

- **US1 (P1)**: Foundational only. No cross-story dependencies
- **US2 (P2)**: Extends location.go from US1. Must complete after US1
- **US3 (P3)**: Extends display.go from US1. Can run parallel with US2 after US1
- **US4 (P3)**: Extends display.go from US1. Can run parallel with US2/US3 after US1

### Within Each User Story

- Types/models before services
- Services before API clients
- API clients before orchestration (main.go)
- Core implementation before tests
- Tests verify implementation

### Parallel Opportunities

- Phase 2: T006, T007, T008, T009, T010 can all run in parallel
- Phase 3: T011 + T012 in parallel (header + implementation), T017 in parallel with API work
- Phase 5 + Phase 6 can run in parallel after Phase 3 completion

---

## Parallel Example: User Story 1

```bash
# Launch foundational tasks in parallel:
Task: "Define CurrentWeather and DailyForecast structs in internal/weather/weather.go"
Task: "Implement WMO weather code mapping in internal/display/conditions.go"
Task: "Implement ANSI color helpers in internal/display/color.go"
Task: "Create test fixtures in testdata/"

# Launch CoreLocation bridge files in parallel:
Task: "CoreLocation C header in internal/location/corelocation_darwin.h"
Task: "CoreLocation Obj-C implementation in internal/location/corelocation_darwin.m"
Task: "ASCII art icons in internal/display/ascii.go"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup (T001-T004)
2. Complete Phase 2: Foundational (T005-T010)
3. Complete Phase 3: User Story 1 (T011-T022)
4. **STOP and VALIDATE**: Build and run `goweather` with no args
5. Verify weather card displays with location, current conditions, and forecast

### Incremental Delivery

1. Setup + Foundational: Project compiles, types defined
2. US1: `goweather` works with auto-location (MVP!)
3. US2: `goweather --city Berlin` works
4. US3: `goweather --imperial` and `--no-color` work
5. US4: `goweather --days 3` works
6. Polish: Edge cases handled, build verified
