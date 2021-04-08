/*
	- Encode structures to the popular JSON data format.
	- Go provides a structure type.
	(Whereas collections are of the same type, structures allow you
	to group disparate things togheter)
*/

package main

import "fmt"

func main() {

	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
}
