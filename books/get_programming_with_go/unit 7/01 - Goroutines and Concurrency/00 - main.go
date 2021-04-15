/*
[Concepts]

	- In Go, an independently running task is known as a goroutine.
	- Goroutines are similar to coroutines, fibers, processes, or threads in other languages,
		although they’re not quite the same as any of those. They’re very efficient to create, and
		Go makes it straightforward to coordinate many concurrent operations.
*/

package main

import (
	"fmt"
	"time"
)

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...")
}

func main() {
	sleepyGopher() //the goroutine is started
	time.Sleep(4 * time.Second)
	fmt.Println("...ronc...")
}
