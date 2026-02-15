# weather

A minimalistic terminal weather CLI for macOS that displays current conditions and a multi-day forecast with colorful ASCII art. Supports six languages and automatic location detection via macOS CoreLocation.

> [!NOTE]
> This project was created as a smoke test for exploring [SDD](https://github.com/rhuss/cc-sdd) (Spec-Driven Development) and its Claude Code plugin. It is not actively maintained and won't receive further updates.

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

## Example Output

```
$ weather -city Berlin -days 3
┌────────────────────────────────────────────────────┐
│  Berlin, Germany  ⛅                               │
│                                                    │
│                  ⛅ Partly cloudy                  │
│       .--.       -3°C (feels -7°C)                 │
│    .-(    ).     Humidity: 72%                     │
│   (___.__)__)    Wind: 5 km/h NNW                  │
│                                                    │
│                                                    │
├────────────────────────────────────────────────────┤
│  Day           Hi    Lo  Cond.                     │
│  Sun 15      -0°C  -3°C  ☁️ Overcast               │
│  Mon 16      -2°C  -4°C  ❄️ Slight snow            │
│  Tue 17       2°C  -2°C  🌨️ Slight snow showers    │
└────────────────────────────────────────────────────┘
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
