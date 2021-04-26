package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	coditem        string `json:"coditem"`
	nomeitem       string `json:"nomeitem"`
	valitem        string `json:"valitem"`
	qtditem        string `json:"qtditem"`
	descontomaximo string `json:"descontomaximo"`
}

func checkError(err error, i int) {
	if err != nil {
		fmt.Printf("Error: #%d", i)
		fmt.Println("Retorno: ", err.Error())
		os.Exit(1)
	}
}

func main() {

	//ler o arquivo
	csvFile, err := os.Open("sample.csv")
	checkError(err, 1)

	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	checkError(err, 2)

	var item Item
	var itens []Item

	for _, rec := range records {
		item.coditem = rec[0]
		item.nomeitem = rec[1]
		item.valitem = rec[2]
		item.qtditem = rec[3]

		itens = append(itens, item)
	}

	//convertendo para json
	// Convert to JSON
	json_data, err := json.Marshal(itens)
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
