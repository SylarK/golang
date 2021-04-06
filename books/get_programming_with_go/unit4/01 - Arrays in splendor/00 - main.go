/*
	Arrays are ordered collections of elementos with a fixed length. You can collect anything you like.
*/

package main

import "fmt"

func main() {
	var planets [8]string // default -> ""
	var numbers [5]uint   // default -> 0

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]
	fmt.Println(earth)
	fmt.Println(len(planets))
	fmt.Println(numbers[0])
}
