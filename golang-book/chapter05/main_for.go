package main

import "fmt"

func main() {
	var i int = 1

	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("///")
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}
}
