package raindrops

import "fmt"

func buildDrops(sound string, factor int) func(s string, number int) string {
	return func(s string, number int) string {
		if number%factor == 0 {
			return s + sound
		}

		return s
	}
}

func Convert(number int) string {
	pling := buildDrops("Pling", 3)
	plang := buildDrops("Plang", 5)
	plong := buildDrops("Plong", 7)

	drops := plong(plang(pling("", number), number), number)
	if drops == "" {
		return fmt.Sprint(number)
	}

	return drops
}
