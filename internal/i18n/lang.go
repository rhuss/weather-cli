package i18n

// Lang holds all translatable strings for a language.
type Lang struct {
	Code              string
	LabelDay          string
	LabelHi           string
	LabelLo           string
	LabelCond         string
	LabelHumidity     string
	LabelWind         string
	LabelFeels        string
	DayAbbreviations  [7]string  // indexed by time.Weekday (Sun=0..Sat=6)
	Cardinals         [16]string // N, NNE, NE, ENE, E, ESE, SE, SSE, S, SSW, SW, WSW, W, WNW, NW, NNW
	Conditions        map[int]string // WMO code -> description
	TipManualLocation string
}

var registry = map[string]*Lang{}

func register(l *Lang) {
	registry[l.Code] = l
}
