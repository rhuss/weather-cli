# weather

A terminal weather CLI for macOS that displays current conditions and a multi-day forecast with colorful ASCII art. Supports six languages and automatic location detection via macOS CoreLocation.

## Build

```bash
make build
```

This produces the `./weather` binary. Requires Go 1.22+ with cgo enabled (uses CoreLocation on macOS).

## Usage

```bash
# Current location (auto-detected)
./weather

# Specify a city
./weather -city "Berlin"

# Coordinates
./weather -lat 48.8566 -lon 2.3522

# Imperial units
./weather -city "New York" -imperial

# Change language
./weather -lang de

# Adjust forecast days (1-7, default 5)
./weather -days 3

# Disable colors
./weather -no-color
```

## Flags

| Flag | Description |
|------|-------------|
| `-city` | City name for weather lookup |
| `-lat`, `-lon` | Latitude and longitude (must be used together) |
| `-imperial` | Use Fahrenheit and mph |
| `-metric` | Use Celsius and km/h (default) |
| `-lang` | Language: `en`, `de`, `es`, `fr`, `it`, `zh` |
| `-days` | Forecast days, 1-7 (default 5) |
| `-no-color` | Disable ANSI color output |

## Supported Languages

- English (`en`, default)
- German (`de`)
- Spanish (`es`)
- French (`fr`)
- Italian (`it`)
- Chinese (`zh`)
