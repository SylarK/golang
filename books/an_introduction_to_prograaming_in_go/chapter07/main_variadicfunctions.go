package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}

	return total
}

func main() {
	valX := add(10, 20, 30, 40, 50)
	xs := []int{1, 2, 3}

	fmt.Println(valX)
	fmt.Println(add(xs...))
}
