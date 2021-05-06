// server statis files.go

package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	CONST_HOST = "localhost"
	CONST_PORT = "8080"
)

type Person struct {
	Name string
	Age  string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{"Lucas Amado", "26"}
	parsedTemplate, _ := template.ParseFiles("templates/02.html")
	err := parsedTemplate.Execute(w, person)

	if err != nil {
		log.Printf("Error occurred while executing the template or writing its output : ", err)
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(CONST_HOST+":"+CONST_PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server : ", err)
	}
}
