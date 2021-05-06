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

func login(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/04.html")
	parsedTemplate.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(CONST_HOST+":"+CONST_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
