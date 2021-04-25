package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct{
	Id 			string 	`json:"Id"`
	Title 	string 	`json:"Title"`
	Desc 		string 	`json:"desc"`
	Content string 	`json:"content"`
}

var Articles []Article

func handleRequests(){
	// criando uma instance para o mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	
	//before -> http.HandleFunc("/", homePage)
	myRouter.HandleFunc("/", homePage)

	//before -> http.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles", returnAllArticles)

	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)

	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

	//before -> log.Fatal(http.ListenAndServe(":8000", nil))
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main(){

	fmt.Println("REST Api - Usando mux")

	//criando nosso "banco de dados"
	
	Articles = []Article{
		{Id: "1", Title: "Title 1", Desc: "Article Description 1", Content:"Article Content 1"},
		{Id: "2", Title: "Title 2", Desc: "Article Description 2", Content:"Article Content 2"},
		{Id: "3", Title: "Title 3", Desc: "Article Description 3", Content:"Article Content 3"},
	}

	handleRequests()
	fmt.Println()
}

//getAll
func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

//home
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//getOne (usando o mux para pegar uma variável)
func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	//fmt.Fprintf(w, "Key : " + key);

	for _, article := range Articles{
		if article.Id == key{
			json.NewEncoder(w).Encode(article)
		}
	}

	fmt.Println("Endpoint Hit: returnSingleArticle")
}

//Criando novos articles
func createNewArticle(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article

	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)

	fmt.Fprintf(w, "%+v", string(reqBody))
}

/*
{
    "Id": "4", 
    "Title": "Title 4", 
    "desc": "Article Description 4", 
    "content": "Content 4" 
}
*/

func deleteArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	id := vars["id"]

	for index, article := range Articles{
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}