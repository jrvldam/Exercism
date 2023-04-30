package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fancyNumber, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}

	number, err := strconv.Atoi(fancyNumber.Value())
	if err != nil {
		return 0
	}

	return number
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fancyNumber := ExtractFancyNumber(fnb)

	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(fancyNumber))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	integer, ok := i.(int)
	if ok {
		return DescribeNumber(float64(integer))
	}

	float, ok := i.(float64)
	if ok {
		return DescribeNumber(float)
	}

	nb, ok := i.(NumberBox)
	if ok {
		return DescribeNumberBox(nb)
	}

	fnb, ok := i.(FancyNumberBox)
	if ok {
		return DescribeFancyNumberBox(fnb)
	}

	return "Return to sender"
}
