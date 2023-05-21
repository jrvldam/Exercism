package pangram

import "strings"

func IsPangram(input string) bool {
	str := strings.ToLower(input)
	alphabet := getFromAToZ()

	for _, char := range str {
		if _, ok := alphabet[char]; ok {
			alphabet[char] = true
		}
	}

	for _, found := range alphabet {
		if !found {
			return false
		}
	}

	return true
}

func getFromAToZ() map[rune]bool {
	var pamgram = map[rune]bool{}
	for _, char := range "abcdefghijklmnopqrstuvwxyz" {
		pamgram[char] = false
	}

	return pamgram
}
