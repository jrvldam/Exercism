// Package weather provide an easy way to get forecast of Goblinocus cities.
package weather

// CurrentCondition get forecast conditions once Forecast function is called
// allowing access the weather outside the package.
var CurrentCondition string
// CurrentLocation get location of forcast once Forecast function is called
// allowing access the location outside the package.
var CurrentLocation string

// Forecast returns information of a Goblinocus city forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
