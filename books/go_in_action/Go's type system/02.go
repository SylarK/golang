package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	yagami := user{"Yagami", "yagami@email.com"}
	yagami.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	yagami.changeEmail("yagami@gmail.com")
	yagami.notify()

	lisa.changeEmail("lisa@gmail.com")
	lisa.notify()
}
