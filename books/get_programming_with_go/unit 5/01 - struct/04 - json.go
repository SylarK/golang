package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Lat, Long float64
	}

	curiosity := location{-4.5895, 137.6896}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println("JSON: ", bytes)
	fmt.Println(string(bytes))

}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
