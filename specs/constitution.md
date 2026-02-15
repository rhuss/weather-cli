# Weather CLI Constitution

## Core Principles

### I. Single Binary, Zero Config

The tool produces a single self-contained binary with no runtime dependencies beyond macOS system frameworks. Users should get useful output with zero arguments. Sensible defaults over configuration.

### II. CLI Contract

Standard Unix conventions apply:
- Normal output to stdout, errors to stderr
- Exit code 0 on success, non-zero on failure
- Flags use Go `flag` package conventions (single dash, e.g. `-city`)
- `--no-color` disables ANSI codes for piping and redirection

### III. Graceful Degradation

When a preferred data source is unavailable, fall back silently to the next option. CoreLocation fails? Use IP geolocation. Never crash on recoverable errors. Show clear, actionable error messages when recovery is not possible.

### IV. Test Coverage

All packages must have unit tests. Tests must be runnable with `make test` (which calls `go test ./...`). Test files live alongside their source files following Go conventions. Focus on behavior, not implementation details.

### V. Simplicity

Prefer the standard library over third-party dependencies. Avoid premature abstraction. Three similar lines of code are better than an unnecessary helper. No feature flags, no configuration files, no plugin systems.

## API and External Services

### External API Usage
- Use free, unauthenticated APIs only (Open-Meteo, ip-api.com)
- All HTTP calls must have timeouts (30 seconds maximum)
- Handle API errors gracefully with user-friendly messages
- No API keys or secrets required

### Data Flow
- Location resolution: CoreLocation -> IP geolocation -> manual input
- Weather data: Open-Meteo API only
- Geocoding (city name to coordinates): Open-Meteo geocoding API

## Code Quality

### Go Standards
- Follow standard Go conventions (`gofmt`, `go vet`)
- Package names: short, lowercase, no underscores
- Error handling: check all errors, wrap with context using `fmt.Errorf`
- No `init()` functions except for unavoidable cgo setup

### Project Structure
```
main.go              # Entry point, flag parsing, orchestration
internal/
  display/           # Terminal rendering, ASCII art, colors
  i18n/              # Internationalization, language packs
  location/          # Location detection (CoreLocation, IP geo)
  units/             # Unit conversion (metric/imperial)
  weather/           # Weather API client, geocoding
```

### Internationalization
- All user-facing strings go through the `i18n` package
- Language packs are Go source files (e.g., `lang_de.go`), not external files
- Supported languages: en, de, es, fr, it, zh
- English is the default; system locale detection is attempted first

## Security

### Input Validation
- Validate all CLI flag values before use
- Latitude: -90 to 90, Longitude: -180 to 180
- Days: 1 to 7
- City names: passed directly to geocoding API (URL-encoded by HTTP client)

### Network Safety
- HTTPS for all external API calls where supported
- No secrets, tokens, or credentials in the binary or source
- No user data stored to disk

## Build and Release

### Build System
- `make build` produces the binary
- `make test` runs all tests
- `make clean` removes build artifacts
- cgo is required (CoreLocation binding on macOS)

### Platform
- macOS only (Apple Silicon and Intel)
- Go 1.22+ required
- No cross-compilation support needed

## Governance

Constitution amendments should be documented with rationale. Update when new patterns emerge or existing standards prove impractical. Keep this document aligned with actual project practices.

**Version**: 1.0.0 | **Ratified**: 2026-02-15 | **Last Amended**: 2026-02-15
