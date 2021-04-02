/*
	Pointers reference a location in memory where a value is
	stored rather the value itself.

	Pointer is represented using the * (asterisk) followed by
	the type of the stored value.

	We use the & operator to find the address of a variable.

*/

package main

import "fmt"

func zero(xPtr *int){
	*xPtr = 10
	//store the int 10 in the memory location xPtr
}

func zero_anotherway(xPtr *int){
	*xPtr = 100
}

func alt(arg int){
	arg += 20
}

func main(){
	x := 5
	y := 1
	z := new(int)
	zero(&x)
	zero_anotherway(z)
	alt(y)
	fmt.Println(x);
	fmt.Println(y);
	fmt.Println(*z);
}