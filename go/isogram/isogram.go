package isogram

import "unicode"

func IsIsogram(word string) bool {
	charStore := map[rune]int{}
	for _, char := range word {
		if unicode.IsLetter(char) {
			lowerChar := unicode.ToLower(char)
			if _, ok := charStore[lowerChar]; ok {
				return false
			}

			charStore[lowerChar] = 1
		}
	}

	return true
}
