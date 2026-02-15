package i18n

func init() {
	register(&Lang{
		Code:          "it",
		LabelDay:      "Giorno",
		LabelHi:       "Max",
		LabelLo:       "Min",
		LabelCond:     "Cond.",
		LabelHumidity: "Umidità:",
		LabelWind:     "Vento:",
		LabelFeels:    "perc.",
		DayAbbreviations: [7]string{
			"Dom", "Lun", "Mar", "Mer", "Gio", "Ven", "Sab",
		},
		Cardinals: [16]string{
			"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE",
			"S", "SSO", "SO", "OSO", "O", "ONO", "NO", "NNO",
		},
		Conditions: map[int]string{
			0: "Cielo sereno", 1: "Prevalentemente sereno", 2: "Parzialmente nuvoloso", 3: "Coperto",
			45: "Nebbia", 48: "Nebbia con brina",
			51: "Pioggerella leggera", 53: "Pioggerella moderata", 55: "Pioggerella intensa",
			56: "Pioggerella gelata leggera", 57: "Pioggerella gelata intensa",
			61: "Pioggia leggera", 63: "Pioggia moderata", 65: "Pioggia forte",
			66: "Pioggia gelata leggera", 67: "Pioggia gelata forte",
			71: "Neve leggera", 73: "Neve moderata", 75: "Neve forte", 77: "Granelli di neve",
			80: "Rovesci leggeri", 81: "Rovesci moderati", 82: "Rovesci violenti",
			85: "Rovesci di neve leggeri", 86: "Rovesci di neve forti",
			95: "Temporale", 96: "Temporale con grandine leggera", 99: "Temporale con grandine forte",
		},
		TipManualLocation: "Suggerimento: Usa --city o --lat/--lon per specificare una posizione manualmente",
	})
}
