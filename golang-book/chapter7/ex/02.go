/*
	Write a function with one variadic parameter
	that finds the greatest number in a list of numbers.
*/

package main

import "fmt"

func greatest(args ...int) int {
	value := 0
	for _, x := range args {
		if x > value {
			value = x
		}
	}
	return value
}

func main() {
	fmt.Println(greatest(10, 5, 6, 90, 6, 3, 6, 5))
	fmt.Println(greatest(5, 6, 9, 80, 56, 36, 901, 568, 956))
	fmt.Println(greatest(1, 1, 2, 3, 3, 3, 4, 4, 5, 6, 8, 2, 1, 35, 6, 1, 1, 1, 2))
}
