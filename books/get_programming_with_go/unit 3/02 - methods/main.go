package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

// K -> C converts ºK to ºC
func (k kelvin) celsius() celsius{
	return celsius(k - 273.15)
}

// C -> K converts ºC to ºK
func (c celsius) kelvin() kelvin{
	return kelvin(c + 273.15)
}

// Celsius converts ºF to ºC

func (f fahrenheit) celsius() celsius{
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func main(){
	var k kelvin = 294.0
	c := k.celsius()

	fmt.Printf("%fºK is %fºC\n",k, c)
}

/*
	Functions with custom types
*/