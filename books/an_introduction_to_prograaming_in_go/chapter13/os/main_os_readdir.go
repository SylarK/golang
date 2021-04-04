package main

import (
	"fmt"
	"os"
)

func main(){
	dir, err := os.Open(".")
	if err != nil{
		fmt.Println("Something is wrong [1]")
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil{
		fmt.Println("Something is wrong [2]")
		return
	}

	for _, fi := range fileInfos{
		fmt.Println(fi.Name())
	}
}