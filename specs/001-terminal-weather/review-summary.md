# Review Summary: Terminal Weather Display

**Spec:** specs/001-terminal-weather/spec.md | **Plan:** specs/001-terminal-weather/plan.md
**Generated:** 2026-02-14

> Distilled decision points for reviewers. See full spec/plan for details.

---

## Feature Overview

A Go CLI tool (`goweather`) that displays current weather and a multi-day forecast in the terminal using ASCII art and emoji. It auto-detects location via macOS CoreLocation with IP geolocation fallback, fetches weather from Open-Meteo's free API, and renders a compact bordered card. Supports unit preferences (metric/imperial), color toggle, city override, and configurable forecast range.

## Scope Boundaries

- **In scope:** Current weather display, 1-7 day forecast, auto/manual location, metric/imperial units, color toggle, ASCII art icons, emoji annotations
- **Out of scope:** Windows/Linux, offline mode, caching, config files, hourly forecasts, weather alerts, TUI mode, i18n
- **Why these boundaries:** Focused macOS CLI tool. CoreLocation is macOS-specific by design. Keeping scope tight for a single-binary utility.

## Critical Decisions

### CoreLocation via cgo (not a helper binary)
- **Choice:** Direct cgo bridge to CoreLocation framework using Objective-C
- **Alternatives:** Swift helper binary shelling out, IP-only (no native location)
- **Trade-off:** Ties binary to macOS and cgo (no cross-compilation), but gives best location accuracy and single-binary distribution
- **Feedback:** Is cgo complexity justified for a weather tool, or would IP-only be sufficient?

### Open-Meteo API (no API key)
- **Choice:** Open-Meteo for weather data (free, no registration)
- **Alternatives:** OpenWeatherMap (requires key), wttr.in (less control)
- **Trade-off:** No API key friction, but depends on a free service with no SLA
- **Feedback:** Acceptable dependency on a free service for a personal CLI tool?

### Server-side unit conversion
- **Choice:** Request units from Open-Meteo via query params rather than converting client-side
- **Alternatives:** Client-side conversion with a units package
- **Trade-off:** Simpler code, fewer rounding issues, but adds API coupling
- **Feedback:** Is this the right trade-off for simplicity?

## Areas of Potential Disagreement

### ip-api.com as fallback (HTTP, not HTTPS)
- **Decision:** Use ip-api.com free tier (HTTP only on free plan)
- **Why this might be controversial:** HTTP endpoint exposes user's IP geolocation request to network observers
- **Alternative view:** Use ipapi.co (HTTPS on free tier) or accept no fallback
- **Seeking input on:** Is HTTP acceptable for a location-only query in a CLI tool, or should we use an HTTPS alternative?

### macOS-only scope
- **Decision:** Target macOS exclusively due to CoreLocation dependency
- **Why this might be controversial:** Limits audience significantly
- **Alternative view:** Use IP geolocation only, making it cross-platform
- **Seeking input on:** Is the CoreLocation accuracy worth the platform lock-in?

## Naming Decisions

| Item | Name | Context |
|------|------|---------|
| Binary | `goweather` | CLI command name |
| Module | `goweather` | Go module name |
| Packages | `location`, `weather`, `display`, `units` | Internal packages |

## Schema Definitions

### Open-Meteo Response (key fields)
- `current.temperature_2m`, `current.apparent_temperature`, `current.relative_humidity_2m`
- `current.wind_speed_10m`, `current.wind_direction_10m`, `current.weather_code`
- `daily.time[]`, `daily.temperature_2m_max[]`, `daily.temperature_2m_min[]`, `daily.weather_code[]`

## Architecture Choices

- **Pattern:** Single binary CLI with internal packages, cgo bridge for native APIs
- **Components:** 4 internal packages (location, weather, display, units) + main.go orchestrator
- **Integration:** CoreLocation (cgo), Open-Meteo API (HTTP), ip-api.com (HTTP fallback)

## Open Questions

- [ ] Should `--no-color` also strip emoji for maximum pipe compatibility?
- [ ] Should the tool auto-detect terminal width and adjust card layout?

## Risk Areas

| Risk | Impact | Mitigation |
|------|--------|------------|
| Open-Meteo API goes down or adds rate limits | High - tool becomes non-functional | Could add a secondary weather API as fallback |
| CoreLocation permission denied on first run | Med - confusing UX without explanation | Silent fallback to IP geolocation works |
| ip-api.com rate limit (45/min) hit | Low - unlikely for CLI tool | Single request per invocation |
| cgo build complexity | Med - harder to build/distribute | Makefile with correct flags, documented in quickstart.md |

---
*Share this with reviewers. Full context in linked spec and plan.*
