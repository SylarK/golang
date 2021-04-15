package main

import "fmt"

func main() {

	var nowint *int
	var nowstring *string

	fmt.Println(nowint)
	fmt.Println(nowstring)

	value := 100
	nowint = &value

	name := "Calel"
	nowstring = &name

	if nowint != nil && nowstring != nil {
		fmt.Println(*nowint)
		fmt.Println(*nowstring)
	}

}
