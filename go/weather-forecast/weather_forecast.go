//Package weather is a package that forecast current weather conditions of various cities.
package weather

//CurrentCondition is a variable that contains current condition in string.
var CurrentCondition string
//CurrentLocation is a variable that contains the current location as a string.
var CurrentLocation string

// Forecast function that accepts two variables as strings and returns forecast for the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
