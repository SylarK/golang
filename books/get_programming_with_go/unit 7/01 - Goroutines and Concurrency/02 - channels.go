/*
	A chanel can be used to send values safely fron one goroutine to another.
	Like any other Go type, channels can be used as variables, passed to functions, stored in
	a structure, and do almost anything else you want them to do.

	to create a channel use make

	c := make(chan int)

	to receive a value from a channel:

	c <- 99
	r := <-c
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c //receive a value from a channel
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}

func sleepGopher(id int, c chan int) {
	time.Sleep((3 * time.Second))
	fmt.Println("... ", id, " snore")
	c <- id
}
