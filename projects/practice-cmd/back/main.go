package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Data struct{
	usuario string
	senha string
	database string
	origem string
	destino string
	porta string
}

func main(){

	var option int
	var err error
	
	dataFirebird := Data{
		"SPK", "-", "BUSINESS", "D:\\Bases\\HERDORA\\BUSINESS.FDB", "D:\\Bases\\Backups_pg", "3050", 
	}

	dataPostgres := Data{
		"postgres", "-", "Go", "-", "D:\\Bases\\Backups_pg\\backup_bot_razem.backup", "5432",
	} 

	//showData(&data)	
	msgScreen("RAZEM")
	chooseOption(&option)
	

	switch option{
	case 1: 
		showDataFirebird(&dataFirebird)
		command := exec.Command("gbak.exe", "-b", "-user" , dataFirebird.usuario ,"-pas", dataFirebird.senha ,dataFirebird.origem, "D:\\Bases\\Backups_pg\\backup_bot_razem.FBK" )
		//command.Dir = "C:\\Program Files (x86)\\Firebird\\Firebird_3_0\\"
		var out bytes.Buffer
		var stderr bytes.Buffer
		command.Stdout = &out
		command.Stderr = &stderr
		err = command.Run()
		if err != nil{
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    	return
		}else{
			fmt.Println("\nComando executado com sucesso. Aguarde o backup ficar pronto.\nCaminho do arquivo : ",dataFirebird.destino)
		} 
	case 2:
		showDataPostgres(&dataPostgres)
		command := exec.Command("pg_dump.exe", "--host", "localhost", "--port", "5432", "--username", dataPostgres.usuario, "--format", "tar", "--file", dataPostgres.destino, dataPostgres.database)
		//command.Dir = "C:\\PostgreSQL\\12\\bin\\"
		var out bytes.Buffer
		var stderr bytes.Buffer
		command.Stdout = &out
		command.Stderr = &stderr
		err = command.Run()
		if err != nil{
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    	return
		}else{
			fmt.Println("\nComando executado com sucesso. Aguarde o backup ficar pronto.\nCaminho do arquivo : ",dataFirebird.destino)
		}
	case 3:
		fmt.Print("Caminho do origem do arquivo a ser restaurado: ")
		fmt.Scanln(&dataFirebird.origem)
			dataFirebird.origem = strings.ReplaceAll(dataFirebird.origem, "\\", "\\\\")
		fmt.Print("Caminho do destino do arquivo a ser restaurado: ")
		fmt.Scan(&dataFirebird.destino)
		dataFirebird.destino = strings.ReplaceAll(dataFirebird.destino, "\\", "\\\\")
		
		//command := exec.Command("gbak.exe", "-user", dataFirebird.usuario, "-pas", dataFirebird.senha, "-r", "-p",dataFirebird.porta, "-o", dataFirebird.origem, dataFirebird.destino )
		command := exec.Command("gbak.exe", "-r", "-user", dataFirebird.usuario, "-password", dataFirebird.senha, dataFirebird.origem, dataFirebird.destino )
		fmt.Println(command)
		var out bytes.Buffer
		var stderr bytes.Buffer
		command.Stdout = &out
		command.Stderr = &stderr
		err := command.Run()
		if err != nil{
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    	return
		}else{
			fmt.Println("\nComando executado com sucesso. O seu restore ficou pronto.\nCaminho do arquivo : ",dataFirebird.destino)
		}

	case 4:
	default: 
		fmt.Println("Não foi escolhida uma opção válida. Sistema será encerrado.")
		os.Exit(1)
	} 
}

func chooseOption(i *int){

		fmt.Println("\n1 -\tBackup em Firebird\n2 - \tBackup em Postgres\n3 - \tRestore em Firebird\n4 - \tRestore em Postgres")
		fmt.Print("\nDigite a opção escolhida : ")
		fmt.Scanln(i)
		fmt.Println()

}

func getData(data *Data){
	var temp string
	fmt.Print("\nUser : ")
	fmt.Scan(&data.usuario)
	fmt.Print("\nDatabase : ")
	fmt.Scan(&data.database)
	fmt.Print("\nOrigem : ")
	fmt.Scan(&temp)
	if temp != ""{
		data.origem = strings.ReplaceAll(temp, "\\", "\\\\")
	}
	fmt.Print("\nDestino : ")
	fmt.Scan(&temp)
	if temp != ""{
		data.destino = strings.ReplaceAll(temp, "\\", "\\\\")
	}
	
}

func msgScreen(msg string){
	fmt.Println(strings.Repeat("#", 10), msg, strings.Repeat("#", 10))
}

func showDataFirebird(data *Data){
	msgScreen("DADOS ATUAIS")
	var op string
	
	fmt.Printf("Usuario:\t%s\nNome Banco:\t%s\nOrigem Banco:\t%s\nDestino Banco:\t%s\n\n", data.usuario, data.database, data.origem, data.destino)
	fmt.Print("Você deseja alterar essas configurações ? (S/N) ")
	fmt.Scanln(&op)
	op = strings.ToUpper(op)
	if op == string('S'){
		getData(data)
	}
	fmt.Println()
}

func showDataPostgres(data *Data){
	msgScreen("DADOS ATUAIS")
	var op string
	
	fmt.Printf("Usuario:\t%s\nNome Banco:\t%s\nDestino Banco:\t%s\n\n", data.usuario, data.database, data.destino)
	fmt.Print("Você deseja alterar essas configurações ? (S/N) ")
	fmt.Scanln(&op)
	op = strings.ToUpper(op)
	if op == string('S'){
		getData(data)
	}
	fmt.Println()
}


