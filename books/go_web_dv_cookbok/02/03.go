package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

const (
	CONST_HOST = "localhost"
	CONST_PORT = "8080"
)

type Person struct {
	Id   string
	Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{"1", "Lucas Amado"}
	parsedTemplate, _ := template.ParseFiles("templates/02.html")
	err := parsedTemplate.Execute(w, person)

	if err != nil {
		log.Printf("Error occurred while executing the template or writing its output : ", err)
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", renderTemplate).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	err := http.ListenAndServe(CONST_HOST+":"+CONST_PORT, router)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
