package main

import (
	"fmt"
)

// iota
const (
	a = iota * 10
	_
	c
	_
	e
	f
)

// deslocamento
/*
110000
001100
000011
*/

func main() {
	fmt.Println(a, c, e, f)

	x := 20
	y := x << 2
	z := x >> 2

	fmt.Printf("\n%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)

}
