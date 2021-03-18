package main

import (
	"fmt"
)

const (
	_ = 2020 + iota
	year2
	year3
	year4
	year5
)

func main() {
	fmt.Printf("%v\t%v\t%v\t%v", year2, year3, year4, year5)
}
