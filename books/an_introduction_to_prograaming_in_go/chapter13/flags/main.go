package main

import (
	"fmt"
	"flag"
	"math/rand"
)

func main(){
	maxp := flag.Int("max", 6, "the max value")
	flag.Parse()
	//Generate a number between 0 and max

	fmt.Println(rand.Intn(*maxp))
}

