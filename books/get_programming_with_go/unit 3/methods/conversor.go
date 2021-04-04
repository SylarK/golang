//challenge

package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

// K -> C
func (k kelvin) celsius() celsius{
	return celsius(k - 273.15)
}

// K -> F
func (k kelvin) fahrenheit() fahrenheit{
	return fahrenheit(((k-273.15) * 9.0 / 5.0) + 32.0)
}

// C -> K
func (c celsius) kelvin() kelvin{
	return kelvin(c + 273.15)
}

// C -> F
func (c celsius) fahrenheit() fahrenheit{
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// F -> C
func (f fahrenheit) celsius() celsius{
	return celsius((f-32) * 5.0 / 9.0)
}

// F -> K
func (f fahrenheit) kelvin() kelvin{
	return kelvin(((f-32.0) * 5.0 / 9.0) + 273.15)
} 

func main(){
	
	var k kelvin = 400.15;
	kelvintocelsius := k.celsius()
	kelvintofahrenheit := k.fahrenheit()

	var c celsius = 127.0
	celsiustokelvin := c.kelvin()
	celsiustofahrenheit := c.fahrenheit()

	var f fahrenheit = 400.15
	fahrenheittocelsius := f.celsius()
	fahrenheittokelvin := f.kelvin()

	//out
	fmt.Printf("%fºK -> %fºC\n", k, kelvintocelsius)
	fmt.Printf("%fºK -> %fºF\n", k, kelvintofahrenheit)
	fmt.Printf("%fºC -> %fºK\n", c, celsiustokelvin)
	fmt.Printf("%fºC -> %fºF\n", c, celsiustofahrenheit)
	fmt.Printf("%fºF -> %fºC\n", f, fahrenheittocelsius)
	fmt.Printf("%fºF -> %fºK\n", f, fahrenheittokelvin)
}