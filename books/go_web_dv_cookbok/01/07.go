package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	//CONS_HOST constante host padrão
	CONS_HOST = "localhost"

	//CONS_PORT constante porta padrão
	CONS_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	fmt.Println("Endpoint: Hello World")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Page!")
	fmt.Println("Endpoint: Login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout page")
	fmt.Println("Endpoint: Logout")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	err := http.ListenAndServe(CONS_HOST+":"+CONS_PORT, nil)
	if err != nil {
		log.Fatal("Erro ao iniciar http server : ", err)
	}
}
