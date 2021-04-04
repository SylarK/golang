/*
	Defer, schedules a function call to be run after the function completes.
	Is often used where resources need to be freed in some way.

	Ex:
	f, _ := os.Open(filename)
	defer f.Close()
*/
package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func main() {
	defer second()
	first()
}
