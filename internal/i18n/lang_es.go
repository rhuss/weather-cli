package i18n

func init() {
	register(&Lang{
		Code:          "es",
		LabelDay:      "Día",
		LabelHi:       "Máx",
		LabelLo:       "Mín",
		LabelCond:     "Cond.",
		LabelHumidity: "Humedad:",
		LabelWind:     "Viento:",
		LabelFeels:    "sens.",
		DayAbbreviations: [7]string{
			"Dom", "Lun", "Mar", "Mié", "Jue", "Vie", "Sáb",
		},
		Cardinals: [16]string{
			"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE",
			"S", "SSO", "SO", "OSO", "O", "ONO", "NO", "NNO",
		},
		Conditions: map[int]string{
			0: "Cielo despejado", 1: "Mayormente despejado", 2: "Parcialmente nublado", 3: "Nublado",
			45: "Niebla", 48: "Niebla con escarcha",
			51: "Llovizna ligera", 53: "Llovizna moderada", 55: "Llovizna intensa",
			56: "Llovizna helada ligera", 57: "Llovizna helada intensa",
			61: "Lluvia ligera", 63: "Lluvia moderada", 65: "Lluvia intensa",
			66: "Lluvia helada ligera", 67: "Lluvia helada intensa",
			71: "Nevada ligera", 73: "Nevada moderada", 75: "Nevada intensa", 77: "Granizo fino",
			80: "Chubascos ligeros", 81: "Chubascos moderados", 82: "Chubascos violentos",
			85: "Chubascos de nieve ligeros", 86: "Chubascos de nieve intensos",
			95: "Tormenta", 96: "Tormenta con granizo ligero", 99: "Tormenta con granizo intenso",
		},
		TipManualLocation: "Consejo: Use --city o --lat/--lon para especificar una ubicación manualmente",
	})
}
