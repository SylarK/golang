/*
	If you need multiple structures with the same fields, you can
	define a type.
*/
package main

import "fmt"

func main() {

	type location struct {
		lat  float64
		long float64
	}

	//without composite literals
	var spirit location

	spirit.lat = -14.5600
	spirit.long = 175.47566

	//with composite literals
	oppor := location{lat: -1.9598, long: 354.6596}

	//composite literals without specify field names
	curiosity := location{-4.6598, 138.655}

	fmt.Println(spirit, oppor, curiosity)

	fmt.Printf("%v\n", curiosity)
	//{-4.6598 138.655}
	fmt.Printf("%+v\n", curiosity)
	//{lat:-4.6598 long:138.655}
}
