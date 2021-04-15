package main

import "fmt"

type person struct {
	name       string
	superpower string
	age        int
}

func (p *person) birthday() {
	p.age++
	if p.age%2 == 0 {
		p.superpower = "No superpower"
	} else {
		p.superpower = "Fire"
	}
}

func main() {

	person := person{"Max", "Fire", 25}

	fmt.Printf("%v\t%T\n", person, person) // {Max Fire 25}  main.person
	person.birthday()                      // +1

	fmt.Printf("%v\t%T\n", person, person) // {Max No superpower 26}  main.person
	person.birthday()                      // +1

	fmt.Printf("%v\t%T\n", person, person) // {Max Fire 27}   main.person

}
