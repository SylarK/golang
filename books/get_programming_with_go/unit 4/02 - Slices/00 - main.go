/*
	Slicing doesn't alter one array. It just creates a window or view into the array.
	This view is a type called a slice.
	If you have a collection, is it organized in a certain way? The books on a
	library shelf may be ordered by the last name of the author, for example.
	This arrangement allows you to focus in on other books they wrote
*/

package main

import "fmt"

func main() {

	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	giants := planets[4:] // Ommiting the last index defaults to the lenth of the array
	gas := giants[:2]     // Ommiting the first index defaults to the beggining of the array
	ice := giants[2:4]

	fmt.Println(terrestrial, gasGiants, iceGiants)
	fmt.Println(giants, gas, ice)

	iceGiantsMarkII := iceGiants
	iceGiants[1] = "Poseidon"
	fmt.Println(planets)
	fmt.Println(iceGiants, iceGiantsMarkII, ice)
}
