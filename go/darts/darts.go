package darts

import "math"

const (
	innerCircle  = 1
	middleCircle = 5
	outerCircle  = 10
)

func Score(x, y float64) int {
	switch {
	case isSideCircle(innerCircle, x, y):
		return 10
	case isSideCircle(middleCircle, x, y):
		return 5
	case isSideCircle(outerCircle, x, y):
		return 1
	default:
		return 0
	}
}

func isSideCircle(r float64, x float64, y float64) bool {
	return math.Pow(0-x, 2)+math.Pow(0-y, 2) <= math.Pow(r, 2)
}

