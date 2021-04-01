/*
	Panic & Recover

	recover -> stops the panic and returns the value that was passed to the call to panic.

	panic -> run time error
	A panic generally indicates a programmer error (for
	example attempting to access an index of an array
	that's out of bounds, forgetting to initialize a map, etc.)
	or an exceptional condition that there's no easy way to
	recover from. (Hence the name “panic”)
*/
package main

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

}
