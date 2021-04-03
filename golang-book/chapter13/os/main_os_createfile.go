package main

import (
	"os"
)

func main(){
	file, err := os.Create("secondtest.txt")
	if err != nil{
		return
	}
	defer file.Close()
	
	file.WriteString("testing...")
}