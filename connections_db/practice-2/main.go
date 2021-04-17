package main

import (
	"database/sql"
	"fmt"

	dbConfig "dbconfig"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
var option int
var table string = "product"

type Usuario struct {
	id, age                      int
	first_name, last_name, email string
}

const (
	host   = "localhost"
	dbname = "Go"
	user   = "postgres"
	pass   = "masterkey"
	port   = 5432
)

// criando uma estrutura para produto

func checkErr(err error, i int) {
	if err != nil {
		fmt.Println("*********", i, "*********")
		panic(err.Error())
	}
}

func main() {

	fmt.Printf("Accessing %s ...\n", dbname)

	db, err := sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)
	if err != nil {
		fmt.Println("Erro na conexão")
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

	//sqlInsert()
	//sqlSelectID()
	//sqlUpdate()
	//sqlDelete()
	//sqlSelectAll()

	sqlSelectID()

}

func sqlSelectID() {

	var user Usuario

	sql := fmt.Sprintf("SELECT id, age, first_name, last_name, email FROM users WHERE id=$1")

	err = db.QueryRow(sql, 1).Scan(&user.id, &user.age, &user.first_name, &user.last_name, &user.email)

	fmt.Println(sql)
	fmt.Println(user)

}
