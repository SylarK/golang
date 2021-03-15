package main

import (
	"fmt"
)

	type dino int
	var x dino

func main(){
	fmt.Println(x);
	fmt.Printf("%T\n", x);
	x = 42;
	fmt.Println(x);
}