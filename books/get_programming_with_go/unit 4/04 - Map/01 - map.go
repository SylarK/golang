// map [string] int - map [key type] [value type]

package main

import "fmt"

func main() {

	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	/*
		Go provides the comma, ok syntax, which you can use to distinguish between the "Moon"
		not existing in the map versus being present in the map with a temperature of 0ºC

		--> The moon variable will contain the value found at the "Moon" key or the zero value.
		The additional ok variable will be true if the key is present or false otherwise.

		//Ex. temp, found := temperature["Venus"]
	*/

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %vº C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
}
