package main

import "fmt"


type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) Talk(){
	fmt.Println("Hi, my name is", p.Name)
}

func (a *Android) Info(){
	fmt.Println("My model is", a.Model);
}

func main(){
	a := new(Android);
	a.Name = "B6";
	a.Talk();
	a.Model = "456BoM"
	a.Info();
}