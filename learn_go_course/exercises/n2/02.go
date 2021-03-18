package main

import (
	"fmt"
)

func main() {

	a := 5 == 5
	b := 5 <= 5
	c := 5 > 5
	d := 5
	e := 5 < 5

	fmt.Printf("%t\n%t\n%t\n%d\n%t\n", a, b, c, d, e)

}
