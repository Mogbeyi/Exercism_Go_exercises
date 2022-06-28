// Package weather is a program for them that can forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition tells you the humidity of the weather at that point in time.
var CurrentCondition string

// CurrentLocation reports the location at that point in time.
var CurrentLocation string

// Forecast the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
