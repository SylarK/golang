/*
	Shared values in a Go programs are a bit like a shared phone. If two or more goroutines
	try to use a shared value at the same time, things can go wrong.

	Mutex --> Mutual exclusion.
	Goroutines can use a mutex to exclude each other from doing something at the same time.
	Mutex have two methods : Lock and Unlock.

*/

package main

import "sync"

var mu sync.Mutex

func main() {
	mu.Lock()
	defer mu.Unlock() // unlock the mutex before returning
}
