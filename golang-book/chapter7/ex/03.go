/*
	Using makeEvenGenerator as an example, write a
	makeOddGenerator function that generates odd
	numbers
*/

package main

import "fmt"

var i uint = 0

func makeEvenGenerator() func() uint {

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func makeOddGenerator() func() uint {
	return func() (ret uint) {
		ret = i
		i -= 2
		return
	}
}

func main() {
	i = 0
	nextEven := makeEvenGenerator()
	backEven := makeOddGenerator()

	fmt.Println(nextEven()) //0 - x:2
	fmt.Println(nextEven()) //2 - x:4
	fmt.Println(nextEven()) //4 - x:6
	fmt.Println(backEven()) //6 - x:4
	fmt.Println(backEven()) //4 - x:2
	fmt.Println(backEven()) //2 - x:0
	fmt.Println(i)
}
