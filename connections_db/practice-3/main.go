package main

import (
	"database/sql"
	"fmt"

	dbConfig "dbconfig"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

type Usuarios struct {
	ID, age                    int
	firstName, lastName, email string
}

func checkErr(err error, i int) {
	if err != nil {
		fmt.Printf("Error : #%d", i)
		panic(err.Error())
	}
}

func main() {

	fmt.Printf("Accessing %s... ", dbConfig.DbName)

	//conexão
	db, err = sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()
	sqlSelect()
}

func sqlSelect() {

	tablename := string(dbConfig.TableName[1])

	sqlStatement, err := db.Query("SELECT id, age, first_name, last_name, email FROM " + tablename)
	checkErr(err, 1)

	for sqlStatement.Next() {
		var usuario Usuarios

		err = sqlStatement.Scan(&usuario.ID, &usuario.age, &usuario.firstName, &usuario.lastName, &usuario.email)
		checkErr(err, 2)

		fmt.Printf("Name : %s %s\n", usuario.firstName, usuario.lastName)
	}
}
