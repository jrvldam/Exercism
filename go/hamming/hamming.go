package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Sequences should have same length")
	}

	bs := []rune(b)

	var distance int
	for idx, char := range a {
		if char != bs[idx] {
			distance += 1
		}
	}

	return distance, nil
}
