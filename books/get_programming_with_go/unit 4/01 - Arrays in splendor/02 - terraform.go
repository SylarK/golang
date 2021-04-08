package main

import "fmt"

func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New" + planets[i]
	}
}

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
	terraform(planets)
	fmt.Println(planets)
}

// The terraform function is operation on a copy of the planets array, so the modifications don't affect
// planets in the main function.
