/*
Errors are values.
Don’t just check errors, handle them gracefully.
Don’t panic.
Make the zero value useful.
The bigger the interface, the weaker the abstraction.
interface{} says nothing.
Gofmt’s style is no one’s favorite, yet gofmt is everyone’s favorite.
Documentation is for users.
A little copying is better than a little dependency.
Clear is better than clever.
Concurrency is not parallelism.
Don’t communicate by sharing memory, share memory by communicating.
Channels orchestrate; mutexes serialize.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	files, err := ioutil.ReadDir(`C:\Program Files\firebird\firebird_3_0`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
