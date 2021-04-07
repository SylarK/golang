package main

import (
	"fmt"
	"strings"
)

type worlds []string

func (w worlds) hyperspace() { // This argument is a slice, not an array
	for i := range w {
		w[i] = strings.TrimSpace(w[i])
	}
}

func teste(temp []int) {
	temp[0] = 4
	temp[1] = 5
	temp[2] = 6
	fmt.Println("Varíavel é igual a ", temp)
}

func main() {
	planets := worlds{" Venus  ", "Earth  ", "   Mars"} //strings with spaces
	var temp = []int{1, 2, 3}
	teste(temp)
	planets.hyperspace() // This function will removes the space surroundind worlds
	fmt.Println(strings.Join(planets, ""))
	fmt.Println(temp)
}

/*
	Slices are more versatile than arrays in other ways too. Slices have a length, but unlike
	arrays, the length isn’t part of the type. You can pass a slice of any size to the hyperspace
	function.
	Arrays are rarely used directly. Gophers prefer slices for their versatility, especially
	when passing arguments to functions.
*/
