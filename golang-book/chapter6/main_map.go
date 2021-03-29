package main

import "fmt"

func main() {

	//var x map[string]int
	//x is a map of strings to ints
	x := make(map[string]int)
	x["key"] = 10

	y := make(map[int]int)
	y[1] = 10

	fmt.Println("X: ", x["key"])
	fmt.Println("Y: ", y[1])

	//

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
	}

	if name, ok := elements["Be"]; ok {
		fmt.Println(name, ok)
	}

	//maps are also often used to store general information
	infos := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}

	if el, ok := infos["H"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
