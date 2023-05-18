package resistorcolor

var colors = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Colors should return the list of all colors.
func Colors() []string {
	colorList := make([]string, len(colors))
	for color, value := range colors {
		colorList[value] = color
	}

	return colorList
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	value, ok := colors[color]
	if !ok {
		return -1
	}

	return value
}
