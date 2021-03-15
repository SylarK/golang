package main

import (
	"fmt"
)

const x = 10
const (
	xx = 11
	yy = 12
	uu = 13
)

var y = 10

var c int
var d float64

func main() {
	d = x
	fmt.Println(d)
}

// The constants are undefined until be used
