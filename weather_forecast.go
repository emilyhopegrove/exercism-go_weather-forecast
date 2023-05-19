//write a comment for `package weather` that describes its contents.
// The package comment should introduce the package and provide information relevant to the package as a whole.
package weather
//clarify the usage of the package variables `CurrentCondition` and `CurrentLocation`
//tell any user of the package what information the variables store, and what they can do with it.

//
var CurrentCondition string
var CurrentLocation string
//write a comment for this function that describes what the function does, but not how it does it.
//
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
