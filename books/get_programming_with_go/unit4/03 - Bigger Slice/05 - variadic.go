/*
	Printf and append are variadic function because they accept a variable number of arguments.
	To declare a variadic function, use the ellipsis... with the last parameter.
*/

package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	twoworlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoworlds)
	threeworlds := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", threeworlds...)
	fmt.Println(newPlanets)
}
