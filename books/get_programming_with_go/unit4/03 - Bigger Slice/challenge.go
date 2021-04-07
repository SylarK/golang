package main

import "fmt"

func loops(cicle int, word string) {
	result := make([]string, cicle)
	for i := range result {
		result[i] = word
	}
	fmt.Println(result)
}

func main() {

	var cicle int = 10
	element := "Mars"

	loops(cicle, element)
}
