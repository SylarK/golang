package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "masterkey"
	dbname   = "Go"
)

var age int
var first_name, last_name, email string

func main() {

	fmt.Print("Enter the first name : ")
	fmt.Scanln(&first_name)
	fmt.Print("Enter the last name : ")
	fmt.Scanln(&last_name)
	fmt.Print("Enter the email : ")
	fmt.Scanln(&email)
	fmt.Print("Enter the age: ")
	fmt.Scanln(&age)

	// host, port, user, password, dbname, ssl
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Erro ao conectar.")
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected in database ", dbname)

	sqlStatement := `
		INSERT INTO users (age, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	id := 0
	err = db.QueryRow(sqlStatement, age, email, first_name, last_name).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is: ", id)
}
