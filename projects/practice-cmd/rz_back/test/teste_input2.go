package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	tipo     string
	usuario  string
	senha    string
	database string
	origem   string
	destino  string
	porta    string
}

/*
 Three ways of taking input
   1. fmt.Scanln(&input)
   2. reader.ReadString()
   3. scanner.Scan()

   Here we recommend using bufio.NewScanner
*/

func main() {
	dataFirebird := Data{
		"FB", "SPK", "-", "BUSINESS", "F:\\Bases\\MYBASE\\BUSINESS.FDB", "C:\\BKPRZ_botFB.FBK", "3050",
	}

	/* dataPostgres := Data{
		"PG", "postgres", "-", "Go", "-", "C:\\Spk Sistemas\\BKPRZ_botPG.backup", "5432",
	} */

	getData(&dataFirebird)
	dataFirebird.destino = strings.ReplaceAll(dataFirebird.destino, "\\", "\\\\")
	dataFirebird.origem = strings.ReplaceAll(dataFirebird.origem, "\\", "\\\\")
	fmt.Println(dataFirebird)

	// test
	/*
		in := bufio.NewReader(os.Stdin)
		line, err := in.ReadString('\n')
		if err != nil {
			panic("Erro ao adquirir string")
		}
		fmt.Println(line) */
}

func getData(data *Data) {
	var temp string
	fmt.Print("\nUsuario : ")
	fmt.Scanln(&data.usuario)
	if data.tipo == "FB" {
		fmt.Println("BACKUP FIREBIRD\n")
		fmt.Print("\nCaminho para o arquivo (ex : C:\\Bases\\rzbusiness.FDB) : ")
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')
		data.origem = line
		//fmt.Scanln(&data.origem)
		fmt.Print("\nCaminho onde será colocado o backup (ex : C:\\Backup\\rzbusiness.FBK) : ")
		line, _ = in.ReadString('\n')
		data.destino = line
	} else {
		fmt.Print("\nNome do database no sistema : ")
		fmt.Scanln(&data.database)
		fmt.Print("\nCaminho onde será colocado o backup (ex : C:\\Backup\\rzbusiness.backup) : ")
		fmt.Scanln(&temp)
		data.destino = strings.ReplaceAll(temp, "\\", "\\\\")
		fmt.Print("\nPorta : ")
		fmt.Scanln(&data.porta)
	}

}
