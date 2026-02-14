# Data Model: Terminal Weather Display

**Date**: 2026-02-14
**Feature**: 001-terminal-weather

## Entities

### Location

Represents a resolved geographic position.

| Field       | Type    | Description                                    |
|-------------|---------|------------------------------------------------|
| Latitude    | float64 | Geographic latitude (-90 to 90)                |
| Longitude   | float64 | Geographic longitude (-180 to 180)             |
| City        | string  | Display name (city or region), may be empty     |
| Country     | string  | Country name, may be empty                      |
| Source      | string  | How location was obtained: "corelocation", "ip", "manual" |

### CurrentWeather

Snapshot of current weather conditions at a location.

| Field              | Type    | Description                              |
|--------------------|---------|------------------------------------------|
| Temperature        | float64 | Current temperature (unit depends on request) |
| ApparentTemperature| float64 | "Feels like" temperature                 |
| Humidity           | int     | Relative humidity percentage (0-100)     |
| WindSpeed          | float64 | Wind speed (unit depends on request)     |
| WindDirection      | int     | Wind direction in degrees (0-360)        |
| WeatherCode        | int     | WMO weather interpretation code          |
| Time               | string  | ISO 8601 observation time                |

### DailyForecast

One day's forecast data.

| Field          | Type    | Description                              |
|----------------|---------|------------------------------------------|
| Date           | string  | Date in YYYY-MM-DD format                |
| TemperatureMax | float64 | Daily high temperature                   |
| TemperatureMin | float64 | Daily low temperature                    |
| WeatherCode    | int     | WMO weather interpretation code          |

### WeatherCondition

Static mapping from WMO code to display properties.

| Field       | Type   | Description                                |
|-------------|--------|--------------------------------------------|
| Code        | int    | WMO weather code (0-99)                    |
| Description | string | Human-readable condition ("Clear sky")     |
| Emoji       | string | Emoji representation ("sun")               |
| AsciiArt    | string | Multi-line ASCII art icon                  |
| Category    | string | Grouping: "clear", "cloudy", "rain", "snow", "storm", "fog" |

### Config

Runtime configuration from CLI flags.

| Field      | Type   | Description                                   |
|------------|--------|-----------------------------------------------|
| City       | string | Manual city name (from --city flag)            |
| Latitude   | float64| Manual latitude (from --lat flag), 0 if unset  |
| Longitude  | float64| Manual longitude (from --lon flag), 0 if unset |
| Imperial   | bool   | Use imperial units (from --imperial flag)      |
| NoColor    | bool   | Disable ANSI colors (from --no-color flag)     |
| Days       | int    | Forecast days to show (from --days flag, default 5) |

## Relationships

```
Config → Location resolution strategy
Location → Weather API query (lat/lon)
CurrentWeather + []DailyForecast → Display rendering
WeatherCode → WeatherCondition lookup (static map)
```

## WMO Code Mapping (subset)

| Code | Description          | Category | Emoji |
|------|----------------------|----------|-------|
| 0    | Clear sky            | clear    | sun    |
| 1    | Mainly clear         | clear    | sun behind small cloud |
| 2    | Partly cloudy        | cloudy   | sun behind cloud |
| 3    | Overcast             | cloudy   | cloud  |
| 45   | Fog                  | fog      | fog    |
| 48   | Depositing rime fog  | fog      | fog    |
| 51   | Light drizzle        | rain     | cloud with rain |
| 53   | Moderate drizzle     | rain     | cloud with rain |
| 55   | Dense drizzle        | rain     | cloud with rain |
| 61   | Slight rain          | rain     | cloud with rain |
| 63   | Moderate rain        | rain     | cloud with rain |
| 65   | Heavy rain           | rain     | cloud with rain |
| 71   | Slight snow          | snow     | snowflake |
| 73   | Moderate snow        | snow     | snowflake |
| 75   | Heavy snow           | snow     | snowflake |
| 80   | Slight rain showers  | rain     | sun behind rain cloud |
| 81   | Moderate rain showers| rain     | sun behind rain cloud |
| 82   | Violent rain showers | rain     | cloud with lightning and rain |
| 85   | Slight snow showers  | snow     | cloud with snow |
| 86   | Heavy snow showers   | snow     | cloud with snow |
| 95   | Thunderstorm         | storm    | cloud with lightning |
| 96   | Thunderstorm + hail  | storm    | cloud with lightning |
| 99   | Thunderstorm + heavy hail | storm | cloud with lightning |
