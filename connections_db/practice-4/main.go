package main

import (
	"database/sql"
	dbSpk "dbspk"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error
var lastval int


type Item struct{
	codempresa			int
	coditem 				int
	qtdmovimento  	int
	valormovimento	float64
}

func checkErr(err error, i int){
	if err != nil {
		fmt.Println("#", i)
		panic(err.Error())
	}
}

func main(){

	fmt.Printf("Acessando %s... ", dbSpk.DbName)

	db, err := sql.Open(dbSpk.PostgresDriver, dbSpk.DataSourceName)

	if err != nil {
		panic(err.Error())
	}else{
		fmt.Println("Conectado!\n")
	}

	defer db.Close()

	var item Item
	getData(&item)
	lastValue(db)
	ano, mes, dia := time.Now().Date();
	insertItem(db, item, int(dia), int(mes), int(ano))
}

func lastValue(db *sql.DB){
	sqlStatement := fmt.Sprintf("SELECT MAX(seqmovimento) FROM spkmesfis");
	err = db.QueryRow(sqlStatement).Scan(&lastval)
	checkErr(err, 1)
}

func insertItem(db *sql.DB, item Item, dia, mes, ano int){

	sqlStatement := fmt.Sprintf("INSERT INTO SPKMESFIS (SEQMOVIMENTO, CODEMPRESA, CODITEM, SEQGRADE, SEQLOTE, SEQSERIE, CODALMOXARIF, DTAMOVIMENTO, NUMDOCUMENTO, CODTIPOMOVIMENTO, CODCENTROCUSTO, QTDMOVIMENTOESTOQUE, VALMOVIMENTOESTOQUE, SEQLANCAMENTORESULTADO, DESOBSERVACAO, USUARIOLANCAMENTO)" + 
	" VALUES ($1, $2, $3, 0, NULL, NULL, 1, '%d-%d-%d 00:00:00', 'Ajuste Bot', 6, NULL, $4, $5, NULL, NULL, NULL);", ano, mes, dia)

	insert, err := db.Prepare(sqlStatement)
	checkErr(err ,2)

	result, err := insert.Exec(lastval + 1,item.codempresa, item.coditem, item.qtdmovimento, item.valormovimento)
	checkErr(err, 3)

	affect, err := result.RowsAffected()
	checkErr(err, 4)

	fmt.Printf("\n%d linha inserida. Comando executado com sucesso.\nSequencia do movimento criado : %d", affect, lastval)
}

func getData(item *Item){
	fmt.Print("[ Digite os dados a serem inseridos ]\n")
	fmt.Print("Código da empresa : ")
	fmt.Scanln(&item.codempresa)
	fmt.Print("Código do item : ")
	fmt.Scanln(&item.coditem)
	fmt.Print("Quantidade do movimento : ")
	fmt.Scanln(&item.qtdmovimento)
	fmt.Print("Valor do movimento : ")
	fmt.Scanln(&item.valormovimento)
}
