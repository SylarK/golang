package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strings"

	dbConfig "dbconfig"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
var option int

type Usuarios struct {
	ID, age                    int
	firstName, lastName, email string
}

func checkErr(err error, i int) {
	if err != nil {
		fmt.Printf("Error : #%d\n", i)
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

	
	loop := true

	for loop == true{
		
		options()		

		switch option{

		case 1: 
			sqlSelect()
		case 2:
			var id int
			fmt.Print("Número do ID : ")
			fmt.Scanln(&id)
			sqlSelectID(id)
		case 3:
			var usuario Usuarios
			var id int
			fmt.Print("Número do ID : ")
			fmt.Scanln(&id)
			fmt.Print("First name : ")
			fmt.Scanln(&usuario.firstName)
			fmt.Print("Last name : ")
			fmt.Scanln(&usuario.lastName)
			fmt.Print("Age : ")
			fmt.Scanln(&usuario.age)
			fmt.Print("Email : ")
			fmt.Scanln(&usuario.email)

			sqlUpdate(usuario, id)

		case 4:
			var id int
			fmt.Print("Número do ID : ")
			fmt.Scanln(&id)

			sqlDelete(id)

		case 5: 
			var usuario Usuarios
			fmt.Print("First name : ")
			fmt.Scanln(&usuario.firstName)
			fmt.Print("Last name : ")
			fmt.Scanln(&usuario.lastName)
			fmt.Print("Age : ")
			fmt.Scanln(&usuario.age)
			fmt.Print("Email : ")
			fmt.Scanln(&usuario.email)

			sqlInsert(usuario)

		case 6: 
			fmt.Println("Exit.")
			loop = false
		default: 	
			fmt.Println("Escolha um número válido.")
						
	}
	}

	

	//sqlSelect()
	//sqlSelectID(3)
	//sqlUpdate(2)
	//sqlDelete(3)
	//sqlInsert()
}

func sqlSelect() {

	tablename := string(dbConfig.TableName[1])

	sqlStatement, err := db.Query("SELECT id, age, first_name, last_name, email FROM " + tablename + " ORDER BY ID")
	checkErr(err, 1)

	for sqlStatement.Next() {
		var usuario Usuarios

		err = sqlStatement.Scan(&usuario.ID, &usuario.age, &usuario.firstName, &usuario.lastName, &usuario.email)
		checkErr(err, 2)

		fmt.Printf("Name : %s %s\n", usuario.firstName, usuario.lastName)
	}
}

func sqlSelectID(i int){

	var usuario Usuarios
	
	tablename := string(dbConfig.TableName[1])

	sqlStatement := fmt.Sprintf("SELECT id, age, first_name, last_name, email FROM %s WHERE id=$1", tablename ); 

	err = db.QueryRow(sqlStatement, i).Scan(&usuario.ID, &usuario.age, &usuario.firstName, &usuario.lastName, &usuario.email)
	checkErr(err, 3)

	fmt.Printf("\nId : %d\nFirst name : %s\nLast name : %s\nAge : %d\nEmail : %s\n",usuario.ID,usuario.firstName, usuario.lastName, usuario.age, usuario.email)
}

func sqlUpdate(usuario Usuarios, i int){

	tablename := string(dbConfig.TableName[1])

	sqlStatement := fmt.Sprintf("UPDATE %s SET AGE=$1, FIRST_NAME=$2, LAST_NAME=$3, EMAIL=$4 WHERE ID=$5", tablename)

	update, err := db.Prepare(sqlStatement)
	checkErr(err, 4)

	result, err := update.Exec(usuario.age, usuario.firstName, usuario.lastName, usuario.email, i)
	checkErr(err, 5)

	affect, err := result.RowsAffected()
	checkErr(err, 6)

	fmt.Println(affect, " row(s) affected\tCommand update")

}

func sqlDelete(i int){

	tablename := string(dbConfig.TableName[1])

	sqlStatement := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tablename)

	delete, err := db.Prepare(sqlStatement)
	checkErr(err, 7)

	result, err := delete.Exec(i)
	checkErr(err, 8)

	affect, err := result.RowsAffected()
	checkErr(err, 9)

	fmt.Println(affect ,"row(s) affected\tCommand delete")
}

func sqlInsert(usuario Usuarios){

	tablename := string(dbConfig.TableName[1])

	sqlStatement := fmt.Sprintf("INSERT INTO %s(AGE, first_name, last_name, email) VALUES($1, $2, $3, $4)", tablename)

	insert, err := db.Prepare(sqlStatement)
	checkErr(err, 10)

	result, err := insert.Exec(usuario.age, usuario.firstName, usuario.lastName, usuario.email)
	checkErr(err, 11)

	affect, err := result.RowsAffected()
	checkErr(err, 12)

	fmt.Println(affect, "row(s) affected\tCommand insert")

}

func clear(){
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
}

func div(){
	fmt.Println("\n", strings.Repeat("#", 30), "\n")
}

func options(){
	div()
	fmt.Print("1\tSelect All\n2\tSelect ID\n3\tUpdate\n4\tDelete\n5\tInsert\n6\tExit\n ")
	div()
	fmt.Print("Escolha uma opção: ")
	fmt.Scanln(&option)
}