package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	type location struct {
		Lat  float64 `json:"latitude" xml:"latitude"`
		Long float64 `json:"longitude" xml:"longitude"`
	}
	//the struct tags are formatted as key : "value"

	curiosity := location{-15.3569, 158.2369}

	bytes, err := json.Marshal(curiosity)

	getError(err)

	fmt.Println(string(bytes))
}

func getError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
