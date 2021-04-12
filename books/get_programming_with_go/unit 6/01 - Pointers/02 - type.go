package main

import "fmt"

func main() {
	answer := 42
	name := "gopher"
	address := &answer
	addressGopher := &name
	fmt.Printf("address is a %T\n", address)
	fmt.Printf("address_gopher is a %T\n", addressGopher)
}
