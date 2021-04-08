package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
	fmt.Println(sort.StringSlice(planets).Len())             // 8
	fmt.Println(sort.StringSlice(planets).Less(1, 3))        // true
	fmt.Println(sort.StringSlice(planets).Search("Jupiter")) // 1
}
