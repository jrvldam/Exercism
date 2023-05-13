// Package leap helps to deal with leap years
package leap

// IsLeapYear returns if given year is leap or not
func IsLeapYear(year int) bool {
	if year%100 != 0 && year%4 == 0 {
		return true
	}

	return year%400 == 0
}
