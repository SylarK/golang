/*
It is possible to create functions inside of functions.
When you create a local funcion like this it also has
access to other local variables (scope).
*/

package main

import "fmt"

func main() {

	// 1
	add := func(args ...int) int {
		total := 0
		for _, v := range args {
			total += v
		}
		return total
	}

	fmt.Println("Add: ", add(10, 10, 12, 65, 98, 66, 32, 56, 5))

	// 2
	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	// 3
	nextEven := makeEvenGenerator()

	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
