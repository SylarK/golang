/*
	Write your own Sleep function using time.After
*/

package main

import (
	"fmt"
	"time"
)

func x(arg uint){

	if(arg == 0){
		arg = 1
	}
	for {
		select{
		case <- time.After(time.Second * time.Duration(arg)):
			fmt.Println("Second")
		}
	}
}


func main(){

	x(2)

	var input string 
	fmt.Scanln(&input)
}