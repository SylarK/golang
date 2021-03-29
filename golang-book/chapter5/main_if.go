package main

import "fmt"

func main() {
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			// even
			fmt.Println(j, "even")
		} else {
			// odd
			fmt.Println(j, "odd")
		}
	}
}
