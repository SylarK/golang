package main

import (
	"fmt"
)

type hotdog int
var b hotdog = 101

func main(){
	a := 10
	fmt.Printf("%T", b);
	fmt.Printf("\n%T", a);
	a = int(b)
	fmt.Printf("\nNew value A : %v", a);
}