# Feature Specification: Terminal Weather Display

**Feature Branch**: `001-terminal-weather`
**Created**: 2026-02-14
**Status**: Implemented
**Input**: User description: "Go CLI tool for terminal weather display with emoji and ASCII art, using macOS CoreLocation with IP geolocation fallback and Open-Meteo API"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Quick Weather Check (Priority: P1)

A user opens their terminal and runs `goweather` with no arguments. The tool automatically detects their location and displays current weather conditions along with a 5-day forecast in a compact, visually appealing card with ASCII art and emoji.

**Why this priority**: This is the core value proposition. A zero-configuration weather check is the primary use case.

**Independent Test**: Can be fully tested by running the binary and verifying it outputs a formatted weather card with current conditions and forecast data for the detected location.

**Acceptance Scenarios**:

1. **Given** the user has network connectivity and macOS Location Services enabled, **When** they run `goweather`, **Then** the tool displays a bordered card showing current temperature, conditions with emoji, humidity, wind, and a 5-day forecast table for their current location.
2. **Given** the user has network connectivity but Location Services are disabled, **When** they run `goweather`, **Then** the tool silently falls back to IP-based geolocation and displays weather for the approximate location.
3. **Given** the user has no network connectivity, **When** they run `goweather`, **Then** the tool displays a clear "No network connection" error message.

---

### User Story 2 - Manual Location Override (Priority: P2)

A user wants weather for a specific city or coordinates rather than their current location. They pass `--city Berlin` or `--lat 52.52 --lon 13.41` to get weather for that location.

**Why this priority**: Location override is essential for checking weather at destinations or when automatic detection is inaccurate.

**Independent Test**: Can be tested by running `goweather --city Berlin` and verifying the output shows weather data for Berlin, Germany.

**Acceptance Scenarios**:

1. **Given** the user provides a valid city name, **When** they run `goweather --city "New York"`, **Then** the tool displays weather for New York with the city name in the header.
2. **Given** the user provides latitude and longitude, **When** they run `goweather --lat 48.85 --lon 2.35`, **Then** the tool displays weather for those coordinates.
3. **Given** the user provides an invalid city name, **When** they run `goweather --city "Xyzzyville"`, **Then** the tool displays a "City not found" error message.

---

### User Story 3 - Unit and Display Preferences (Priority: P3)

A user prefers imperial units or needs to pipe the output to another tool. They use `--imperial` for Fahrenheit/mph or `--no-color` for plain text output.

**Why this priority**: Customization options improve usability across different user preferences and tool integration scenarios.

**Independent Test**: Can be tested by running `goweather --imperial` and verifying temperatures are in Fahrenheit, or `goweather --no-color` and verifying no ANSI escape codes are present in the output.

**Acceptance Scenarios**:

1. **Given** default settings, **When** the user runs `goweather`, **Then** temperatures are displayed in Celsius and wind speed in km/h (metric).
2. **Given** the user prefers imperial units, **When** they run `goweather --imperial`, **Then** temperatures are in Fahrenheit and wind speed in mph.
3. **Given** the user wants to pipe output, **When** they run `goweather --no-color`, **Then** the output contains no ANSI color escape codes but retains box-drawing characters and emoji.

---

### User Story 4 - Custom Forecast Range (Priority: P3)

A user wants to see fewer or more forecast days. They use `--days 3` to see a shorter forecast.

**Why this priority**: Minor convenience feature that adds flexibility to the display.

**Independent Test**: Can be tested by running `goweather --days 3` and counting the forecast entries in the output.

**Acceptance Scenarios**:

1. **Given** default settings, **When** the user runs `goweather`, **Then** a 5-day forecast is displayed.
2. **Given** the user wants a shorter forecast, **When** they run `goweather --days 3`, **Then** only 3 days are shown.
3. **Given** the user requests more than 7 days, **When** they run `goweather --days 10`, **Then** the tool displays an error indicating the maximum is 7 days.

---

### User Story 5 - Multi-Language Support (Priority: P3)

A user who prefers a non-English language uses `--lang de` or relies on their system locale to see weather output in their language, including condition descriptions, UI labels, day abbreviations, and cardinal directions.

**Why this priority**: Internationalization broadens the tool's audience beyond English speakers.

**Independent Test**: Can be tested by running `goweather --lang de` and verifying that labels, condition descriptions, and day names appear in German.

**Acceptance Scenarios**:

1. **Given** a user with a German system locale, **When** they run `goweather`, **Then** the output uses German labels, day abbreviations, and weather descriptions.
2. **Given** a user passes `--lang es`, **When** they run `goweather --lang es`, **Then** the output is in Spanish regardless of system locale.
3. **Given** a user passes an unsupported language code, **When** they run `goweather --lang xx`, **Then** the tool falls back to English.

---

### Edge Cases

- What happens when the user is behind a VPN? IP geolocation may return incorrect location. The tool proceeds with the IP-based location; documentation notes users should use `--city` in this case.
- What happens when the terminal is narrower than 40 columns? The tool truncates or wraps content gracefully, maintaining readability.
- What happens when the weather API returns an unknown weather code? The tool displays a generic "Unknown" condition with a question mark emoji.
- What happens when both `--city` and `--lat`/`--lon` are provided? The tool prioritizes `--lat`/`--lon` over `--city` and ignores the city flag.
- What happens when only `--lat` is provided without `--lon` (or vice versa)? The tool displays an error: "Both --lat and --lon must be provided together."

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST detect the user's geographic location automatically using macOS native location services as the primary method.
- **FR-002**: System MUST fall back to IP-based geolocation when native location services are unavailable or access is denied, without displaying an error to the user.
- **FR-003**: System MUST fetch current weather conditions including temperature, "feels like" temperature, humidity percentage, wind speed, wind direction, and weather condition description.
- **FR-004**: System MUST fetch a multi-day forecast (1-7 days, default 5) including daily high and low temperatures and weather condition per day.
- **FR-005**: System MUST display weather data in a compact bordered card using Unicode box-drawing characters.
- **FR-006**: System MUST display an ASCII art icon representing the current weather condition (sun, clouds, rain, snow, thunderstorm, fog, etc.).
- **FR-007**: System MUST annotate weather conditions with appropriate emoji (e.g., sun for clear, cloud for overcast, umbrella for rain, snowflake for snow).
- **FR-008**: System MUST display the location name (city/region) in the card header.
- **FR-009**: System MUST support `--metric` and `--imperial` flags for unit selection, defaulting to metric.
- **FR-010**: System MUST support `--no-color` flag to disable ANSI color codes in output.
- **FR-011**: System MUST support `--city <name>` flag for manual city-based location lookup.
- **FR-012**: System MUST support `--lat <float>` and `--lon <float>` flags for coordinate-based location, requiring both when either is provided.
- **FR-013**: System MUST support `--days <n>` flag to control forecast range (1-7, default 5).
- **FR-014**: System MUST produce a single self-contained binary with no runtime dependencies beyond macOS system frameworks.
- **FR-015**: System MUST support `--lang <code>` flag for language selection (en, de, es, fr, it, zh), with automatic locale detection as the default.
- **FR-016**: System MUST display translated weather condition descriptions, UI labels, day abbreviations, and cardinal directions when a non-English language is active.

### Key Entities

- **Location**: Represents a geographic position with latitude, longitude, and display name (city/region). Can be sourced from CoreLocation, IP geolocation, or user input.
- **CurrentWeather**: Current conditions snapshot including temperature, feels-like temperature, humidity, wind speed, wind direction, and weather condition code.
- **DailyForecast**: A single day's forecast with date, high temperature, low temperature, and weather condition code.
- **WeatherData**: Bundles current conditions with the daily forecast and timezone.
- **WeatherCondition**: A mapping from weather codes to human-readable descriptions, emoji, and category. ASCII art is resolved by category rather than stored per condition.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can view current weather and forecast for their location by running a single command with no arguments.
- **SC-002**: Weather data is displayed within 3 seconds of command execution under normal network conditions.
- **SC-003**: Location is automatically detected without user configuration on a fresh macOS installation with Location Services enabled.
- **SC-004**: Output renders correctly (no garbled characters, proper alignment) in Terminal.app and iTerm2 at 40-column minimum width.
- **SC-005**: Output with `--no-color` flag contains zero ANSI escape sequences and can be redirected to a file without artifacts.
- **SC-006**: All defined weather conditions (clear, cloudy, rain, snow, thunderstorm, fog, drizzle) have distinct ASCII art icons and emoji representations.
- **SC-007**: Output language matches `--lang` flag or system locale, with English as the fallback for unsupported languages.

## Assumptions

- The target platform is macOS (Apple Silicon and Intel).
- Users have an active internet connection for weather data retrieval.
- The Open-Meteo API remains free and available without API keys.
- IP geolocation provides city-level accuracy, which is sufficient for weather purposes.
- Metric units (Celsius, km/h) are the sensible default for an international audience.

## Dependencies

- macOS Location Services (CoreLocation framework)
- Open-Meteo weather API (free, no authentication)
- IP geolocation service (free tier)

## Out of Scope

- Windows or Linux support
- Offline mode or data caching
- Configuration file support
- Hourly forecast breakdown
- Severe weather alerts or warnings
- Interactive or TUI mode
