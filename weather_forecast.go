//Package weather provides the user with the local weather conditions.
package weather

//CurrentCondition represnts a certain weather condition. 
var CurrentCondition string
//CurrentLocation represents a certain location on planet earth.
var CurrentLocation string

//Forecast returns a string description of the current local weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
