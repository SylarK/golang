package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	salary    string `json:"salary"`
}

func checkError(err error, i int) {
	if err != nil {
		fmt.Println("Error: #", i)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func handleRequest() {

	//http.HandleFunc("/", homePage)
	routes := mux.NewRouter().StrictSlash(true)
	routes.HandleFunc("/", homePage)
	routes.HandleFunc("/people", getPeople)
	routes.HandleFunc("/person/{id}", getPerson)

	log.Fatal(http.ListenAndServe(":8000", routes))
}

func main() {
	handleRequest()
	fmt.Println()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: homePage")
	fmt.Fprintf(w, "Home Page")
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: Get People")
	var json_data []Person
	getCSV(&json_data)
	json.NewEncoder(w).Encode(json_data)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	//fmt.Fprintf(w, "Get Person %v\n", key)

	var records []Person
	getCSV(&records)

	for _, rec := range records {
		if rec.ID == key {
			json.NewEncoder(w).Encode(rec)
		}
	}

	fmt.Printf("Endpoint: Get Person {%v}", key)
}

func getCSV(json_data *[]Person) {
	csv_file, err := os.Open("sample.csv")
	checkError(err, 1)
	defer csv_file.Close()

	re := csv.NewReader(csv_file)
	records, err := re.ReadAll()
	checkError(err, 2)

	var person Person
	var people []Person

	for _, rec := range records {
		person.ID = rec[0]
		person.FirstName = rec[1]
		person.LastName = rec[2]
		person.salary = rec[3]

		people = append(people, person)
	}

	*json_data = people
}

func saveJSON() {}
