package main

import "fmt"

func main() {

	answer := 42

	fmt.Println(&answer)
}

//	The address operator (&) provides the memory address of a value.
//	The reverse operation is known as dereferencing which provides the value that a memory address refers to.
//	The following listing deferences the address variable by prefixing it with an asterisk (*)
//	Pointers store memory addresses
