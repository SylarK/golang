/*

Write a program to terraform a slice of strings by prepending each planet with "New ".
Use your program to terraform Mars, Uranus, and Neptune.
Your first iteration can use a terraform function, but your final implementation should
introduce a Planets type with a terraform method, similar to sort.StringSlice

*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

type planets []string

func terraform(p planets, change [3]string) {
	for i := range change {
		for x := range p {
			if p[x] == change[i] {
				p[x] = "New " + p[x]
			}
		}
	}
}

func main() {
	planets := planets{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	change := [3]string{"Mars", "Uranus", "Neptune"}

	sort.StringSlice(planets).Sort()
	terraform(planets, change)
	fmt.Println(planets)
	fmt.Println(strings.Join(planets, ", "))
}
