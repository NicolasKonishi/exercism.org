// Package weather provides weather forecasts for cities in Goblinocus.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string

	// CurrentLocation represents the current location.
	CurrentLocation string
)

// Forecast takes a city and its current weather condition and returns a formatted weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}