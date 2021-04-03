/*
Write a function which takes an integer and
halves it and returns true if it was even or false
if it was odd. For example half(1) should return
(0, false) and half(2) should return (1,
true)
*/

package main

import "fmt"

func verify(arg int) (int, bool) {

	/* value := (arg%2 == 0)

	return arg, value */

	if arg%2 == 0 {
		return 0, true
	} else {
		return 1, false
	}
}

func main() {
	fmt.Println(verify(2))
	fmt.Println(verify(3))
	fmt.Println(verify(4))
	fmt.Println(verify(4))
	fmt.Println(verify(5))
}
