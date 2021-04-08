// Composite literal

package main

import "fmt"

func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "makemake", "Heris"}
	func() {
		for i := 0; i < len(dwarfs); i++ {
			fmt.Println(i, dwarfs[i])
		}
	}()
	fmt.Println("********************")
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	func() {
		for i, planet := range planets {
			fmt.Println(i, planet)
		}
	}()
	//fmt.Println(len(planets))
}
