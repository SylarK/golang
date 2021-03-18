package main

import (
	"fmt"
)

func main() {

	var i int = 200

	fmt.Printf("Bin : %b\nDec : %d\nHex : %#X\n", i, i, i)

	// Deslocamento

	d := i << 1
	fmt.Printf("Bin : %b\nDec : %d\nHex : %#X", d, d, d)

}
