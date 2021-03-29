package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	y := [5]float64{98, 99, 65, 36, 55}

	var total1 float64 = 0

	var total2 float64 = 0

	//1
	for i := 0; i < len(x); i++ {
		total1 += x[i]
	}

	//2
	for _, value := range y {
		total2 += value
	}

	fmt.Println("Average [1]: ", total1/float64(len(x)))
	fmt.Println("Average [2]: ", total2/float64(len(x)))
}
