package main

import "fmt"

type user struct {
	name       string
	email      string
	nvl        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

func main() {

	var natanael user

	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		nvl:        6,
		privileged: true,
	}

	rafael := user{"Rafael", "rafael@email.com", 6, true}

	fred := admin{
		person: user{"fred", "fred@email.com", 8, true},
		level:  "super",
	}

	fmt.Println()
}
