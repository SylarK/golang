package main

import (
	"log"
	"net/http"
	"text/template"
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
	parsedTemplate, _ := template.ParseFiles("static-files/templates/01.html")
	err := parsedTemplate.Execute(w, person)
	if err != nil {
		log.Printf("Error ocurred while executing the template or writing its output : ", err)
		return
	}
}

func main() {
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(CONST_HOST+":"+CONST_PORT, nil)
	if err != nil {
		log.Fatal("Error starting http serve : ", err)
		return
	}
}
