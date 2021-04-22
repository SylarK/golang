//basic and concepts

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

//Utilizando essa variável global apenas para simular um banco de dados

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello Guys! 1", Desc: "Description! 1 ", Content: "Content! SHOW!"},
		Article{Title: "Hello Guys! 2", Desc: "Description! 2 ", Content: "Content! TOP!"},
		Article{Title: "Hello Guys! 3", Desc: "Description! 3", Content: "Content! UHUL!"},
		Article{Title: "Hello Guys! 4", Desc: "Description! 4", Content: "Content! COOL!"},
	}

	handleRequests()
}

//GET ALL
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// json.NewEncoder(w).Encode(article) does the job of encoding our articles array into a JSON
// string and then writing as part of our response.
