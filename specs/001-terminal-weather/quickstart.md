# Quickstart: Terminal Weather Display

**Date**: 2026-02-14
**Feature**: 001-terminal-weather

## Prerequisites

- macOS 10.14 or later
- Go 1.22 or later (with cgo support)
- Xcode Command Line Tools (`xcode-select --install`)

## Build

```bash
# Clone and enter the project
cd goweather

# Build the binary (embeds Info.plist for location permissions)
make build

# Or manually:
CGO_LDFLAGS_ALLOW="-sectcreate" \
  go build -ldflags='-extldflags "-sectcreate __TEXT __info_plist ./Info.plist"' \
  -o goweather .
```

## Run

```bash
# Auto-detect location, show weather
./goweather

# Specify a city
./goweather --city Berlin

# Use coordinates
./goweather --lat 48.85 --lon 2.35

# Imperial units
./goweather --imperial

# Shorter forecast
./goweather --days 3

# No colors (for piping)
./goweather --no-color > weather.txt
```

## First Run

On the first run without `--city` or `--lat`/`--lon`, macOS will prompt for location access permission. If denied, the tool silently falls back to IP-based geolocation.

## Test

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./internal/...
```

## Example Output

```
┌─────────────────────────────────┐
│  Berlin, Germany          sun   │
│                                 │
│    \   /     Clear sky          │
│     .-.      18°C (feels 16°C) │
│  ― (   ) ―  Humidity: 55%      │
│     `-'     Wind: 12 km/h W    │
│    /   \                        │
│                                 │
├─────────────────────────────────┤
│  Day        Hi    Lo    Cond.  │
│  Mon 14     18°   12°   sun    │
│  Tue 15     20°   13°   cloud  │
│  Wed 16     17°   10°   rain   │
│  Thu 17     15°    9°   rain   │
│  Fri 18     19°   11°   sun    │
└─────────────────────────────────┘
```
