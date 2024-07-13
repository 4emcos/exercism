// Package weather provides weather forecast based on city and condition
// current weather condition.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string

// CurrentLocation is the current location.
var CurrentLocation string

// Forecast returns weather forecast based on city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
