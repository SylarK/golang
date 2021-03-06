/*
The Fibonacci sequence is defined as: fib(0) =
0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
Write a recursive function which can find
fib(n).
*/

package main

import "fmt"

func fibonacci(arg uint) uint {
	if(arg == 0){
		return 0
	} else if (arg == 1){
		return 1
	} else {
		return fibonacci(arg-1) + fibonacci(arg-2)
	}	
}

func main() {
	fmt.Println(fibonacci(6))
}
