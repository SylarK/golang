package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func checkError(err error, i int) {
	if err != nil {
		fmt.Println("Erro #", i)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	salary    string `json:"salary"`
}

func main() {
	csv_file, err := os.Open("sample.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csv_file.Close()

	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	checkError(err, 1)

	var emp Employee
	var employees []Employee

	for _, rec := range records {
		emp.ID = rec[0]
		emp.FirstName = rec[1]
		emp.LastName = rec[2]
		emp.salary = rec[3]

		employees = append(employees, emp)
	}
	// Convert to JSON
	json_data, err := json.Marshal(employees)
	checkError(err, 2)

	//print json data
	fmt.Println(string(json_data))

	//create json file
	json_file, err := os.Create("sample.json")
	checkError(err, 3)

	defer json_file.Close()

	json_file.Write(json_data)
	json_file.Close()
}
