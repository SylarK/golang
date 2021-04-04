//SHA-1

package main

import (
	"fmt"
	"crypto/sha1"
)

func main(){
	h := sha1.New()
	h.Write([]byte("Noak"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}