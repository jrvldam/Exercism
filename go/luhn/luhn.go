package luhn

import (
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`^([0-9]|\s)*$`)

func Valid(id string) bool {
	if !regex.MatchString(id) {
		return false
	}

	reversedStr := reverse(strings.ReplaceAll(id, " ", ""))
	if len(reversedStr) < 2 {
		return false
	}

	var sum int
	for idx, char := range reversedStr {
		num := int(char - '0')
		if idx%2 != 0 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		sum += num
	}

	return sum%10 == 0
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
