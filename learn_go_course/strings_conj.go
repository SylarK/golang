package main

import (
	"fmt"
)

func main() {
	s := "Hello, everybody!"
	sb := []byte(s)

	fmt.Printf("%v\n%T\n", s, s)
	fmt.Printf("%v\n%T\n--------\n--------\n", sb, sb)

	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}

}
