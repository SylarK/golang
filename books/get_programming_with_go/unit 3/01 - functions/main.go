//challenge

package main

import "fmt"

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k float64) float64{
	k -= 273.15
	return k
}

// celsiusToFahrenheit converts ºC to ºF
func celsiusToFahrenheit(c float64) float64{
	return (c * 9.0 / 5.0) + 32.0
}

// kelvinToFahrenheit converts ºK to ºF
func kelvinToFahrenheit(k float64) float64{
	if k == 0{
		return -459.67
	}else{
		return ((k-273.15) * 9.0 / 5.0) + 32.0;
	}
}

func main(){

	kelvin := 294.0;
	celsius := kelvinToCelsius(kelvin)
	fahrenheit := celsiusToFahrenheit(celsius)
	ktof := kelvinToFahrenheit(kelvin)
	fmt.Printf("%f Kelvin = %f Celsius\n", kelvin, celsius)
	fmt.Printf("%f Celsius = %f Fahrenheit\n", celsius, fahrenheit)
	fmt.Printf("%f Kelvin = %f Fahrenheit\n", kelvin, ktof)

}

/*

Func kelvinToCelsius accepts one parameter and returns
one result.
The variable kelvin is completely independent from any other functions.

*/