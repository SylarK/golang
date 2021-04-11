/*
	Interfaces are concerned with what a type can do, not the value it holds.
	Methods express the behavior a type provides, interfaces are declared with
	a set of methods that a type must satisfy.
*/

package main

import (
	"fmt"
	"strings"
)

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

/* type t interface {
	talk() string
} */

func main() {

	var t interface {
		talk() string
	}

	t = martian{}
	fmt.Println(t.talk())

	t = laser(2)
	fmt.Println(t.talk())
}
