package display

// AsciiArt returns multi-line ASCII art for a weather category.
func AsciiArt(category string) string {
	switch category {
	case "clear":
		return `    \   /
     .-.
  - (   ) -
     ` + "`" + `-'
    /   \`
	case "cloudy":
		return `
     .--.
  .-(    ).
 (___.__)__)
            `
	case "rain":
		return `     .-.
    (   ).
   (___(__)
    ' ' ' '
   ' ' ' ' `
	case "snow":
		return `     .-.
    (   ).
   (___(__)
    *  *  *
   *  *  * `
	case "storm":
		return `     .-.
    (   ).
   (___(__)
   /_/ /_/
    /  /   `
	case "fog":
		return `
  _ - _ - _
   _ - _ -
  _ - _ - _
            `
	default:
		return `
     .-.
    (   )
     '-'
    ?   ? `
	}
}
