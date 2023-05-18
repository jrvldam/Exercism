package space

import "strings"

type Planet string

var earthYear = map[Planet]float64{
	"mercury": 0.2408467,
	"venus":   0.61519726,
	"earth":   1.0, // 365.25 Earth days, or 31557600 seconds
	"mars":    1.8808158,
	"jupiter": 11.862615,
	"saturn":  29.447498,
	"uranus":  84.016846,
	"neptune": 164.79132,
}

func secondsToYears(seconds float64) float64 {
	return seconds / (60 * 60 * 24 * 365.245)
}

func Age(seconds float64, planet Planet) float64 {
	p := strings.ToLower(string(planet))
	factor, ok := earthYear[Planet(p)]
	if !ok {
		return -1
	}

	return secondsToYears(seconds) / factor
}
