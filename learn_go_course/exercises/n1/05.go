package main

import (
	"fmt"
)

	type dino int
	var x dino
	var y int

func main(){
	fmt.Println(x);
	fmt.Printf("%T\n", x);
	x = 42;
	fmt.Println(x);
	y = int(x)
	fmt.Printf("Value Y: %v, Type Y: %T", y,y)
}