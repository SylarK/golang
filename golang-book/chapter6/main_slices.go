package main

import "fmt"

func main() {

	//1 - var x []float64
	//2 - x := make([]float64, 5)
	//3 - x := make([]float64, 5, 10)
	//4 - arr := [5]float64{1,2,3,4,5}
	//4 - x := arr[0:5]

	//apend
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	//copy
	slice3 := []int{11, 22, 33}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

}
