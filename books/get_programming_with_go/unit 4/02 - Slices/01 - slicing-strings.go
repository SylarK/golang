package main

import "fmt"

func main() {
	string1 := "String"
	string2 := string1[:3]
	string3 := string1[3:]

	fmt.Printf("%s\n%s\t%s", string1, string2, string3)

	//

	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	//dwarfSlice := dwarfArray[:]

	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	fmt.Printf("\n\ndwarfSlice [format] : %T\ndwarfs [format] : %T", dwarfArray, dwarfs)
}
