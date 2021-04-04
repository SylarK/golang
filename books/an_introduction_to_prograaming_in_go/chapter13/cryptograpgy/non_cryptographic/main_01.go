package main

import (
	"fmt"
	"hash/crc32"
)

func main(){
	h := crc32.NewIEEE()
	h.Write([]byte("Test"))
	v := h.Sum32()
	fmt.Println(v)
}

//Once we've written everything we want we call
//Sum32() to return a uint32.