package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso : conversor <valores> <unidade>")
		os.Exit(1)
	}

	unitOrig := os.Args[len(os.Args)-1]
	valOrig := os.Args[1 : len(os.Args)-1]

	var unitDest string

	if unitOrig == "celsius" {
		unitDest = "fahrenheit"
	} else if unitOrig == "quilometros" {
		unitDest = "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!", unitOrig)
		os.Exit(1)
	}

	for i, v := range valOrig {
		valOrig, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s na posição %d não é um número válido! \n", v, i)
			os.Exit(1)
		}

		var valDest float64

		if unitOrig == "celsius" {
			valDest = valOrig*1.8 + 32
		} else {
			valDest = valOrig / 1.60934
		}
		fmt.Printf("%.2f %s = %.2f %s\n", valOrig, unitOrig, valDest, unitDest)
	}
}

// go run conversor.go 32 27 celsius
// 32.00 celsius = 89.60 fahrenheit
// 27.00 celsius = 80.60 fahrenheit
