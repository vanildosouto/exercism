// Package weather get the forecast of a region
package weather

// CurrentCondition condition actual of a region
var CurrentCondition string

// CurrentLocation region
var CurrentLocation string

// Forecast receive region and condition and return a text formatted
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
