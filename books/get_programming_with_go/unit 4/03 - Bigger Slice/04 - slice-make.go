/*
	The make function specifies both the length (0) and capacity(10).
	Up to 10 elements can be appended before the slice runs out of capacity, causing append
	to allocate a new array.
	The capacity argume is optional.const

	Ex. (start with a length and capacity = 0)
	dwarfs := make([]uint, 3)

	--> [0,0,0]

	What is the benefit?
	Preallocatin with make can set an initial capacity, thereby avoiding additional allocations and copies
	to enlarge the underlying array.
*/

package main

import "fmt"

func main() {
	dwarfs := make([]string, 0, 10)

	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println()
}
