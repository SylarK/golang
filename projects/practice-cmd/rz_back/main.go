package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
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

func main() {

	var option int
	var err error

	dataFirebird := Data{
		"FB", "SPK", "-", "BUSINESS", "F:\\Bases\\MYBASE\\BUSINESS.FDB", "C:\\BKPRZ_botFB.FBK", "3050",
	}

	dataPostgres := Data{
		"PG", "postgres", "-", "Go", "-", "C:\\Spk Sistemas\\BKPRZ_botPG.backup", "5432",
	}

	//showData(&data)
	msgScreen("RAZEM")
	chooseOption(&option)

	switch option {
	case 1:
		showDataFirebird(&dataFirebird)
		str1 := dataFirebird.origem
		fmt.Println(str1)
		panic("Error. Pause.")
		command := exec.Command("gbak.exe", "-b", "-user", dataFirebird.usuario, "-pas", dataFirebird.senha, dataFirebird.origem, dataFirebird.destino)
		//command.Dir = "C:\\Program Files (x86)\\Firebird\\Firebird_3_0\\"
		var out bytes.Buffer
		var stderr bytes.Buffer
		command.Stdout = &out
		command.Stderr = &stderr
		err = command.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			fmt.Scanf("h")
			return
		} else {
			fmt.Print("\nComando executado com sucesso. Aguarde o backup ficar pronto.\nCaminho do arquivo : ", dataFirebird.destino)
			fmt.Scanf("h")
		}
	case 2:
		showDataPostgres(&dataPostgres)
		command := exec.Command("pg_dump.exe", "--host", "localhost", "--port", "5432", "--username", dataPostgres.usuario, "--format", "tar", "--file", dataPostgres.destino, dataPostgres.database)
		fmt.Println("COMMAND : ", command)
		panic("STOP HERE")
		//command.Dir = "C:\\PostgreSQL\\12\\bin\\"
		var out bytes.Buffer
		var stderr bytes.Buffer
		command.Stdout = &out
		command.Stderr = &stderr
		err = command.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			fmt.Scanf("h")
			return
		} else {
			fmt.Println("\nComando executado com sucesso. Aguarde o backup ficar pronto.\nCaminho do arquivo : ", dataPostgres.destino)
			fmt.Scanf("h")
		}
	case 3:
		fmt.Print("Caminho origem do arquivo a ser restaurado: ")
		fmt.Scanln(&dataFirebird.origem)
		dataFirebird.origem = strings.ReplaceAll(dataFirebird.origem, "\\", "\\\\")
		fmt.Print("Caminho destino do arquivo a ser restaurado: ")
		fmt.Scan(&dataFirebird.destino)
		dataFirebird.destino = strings.ReplaceAll(dataFirebird.destino, "\\", "\\\\")

		//command := exec.Command("gbak.exe", "-user", dataFirebird.usuario, "-pas", dataFirebird.senha, "-r", "-p",dataFirebird.porta, "-o", dataFirebird.origem, dataFirebird.destino )
		command := exec.Command("gbak.exe", "-r", "-user", dataFirebird.usuario, "-password", dataFirebird.senha, dataFirebird.origem, dataFirebird.destino)
		fmt.Println("\nComando executado com sucesso. Aguarde o seu restore ficar pronto.")
		var out bytes.Buffer
		var stderr bytes.Buffer
		command.Stdout = &out
		command.Stderr = &stderr
		err := command.Run()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			fmt.Scanf("h")
			return
		} else {
			fmt.Println("O seu restore ficou pronto.\nCaminho do arquivo : ", dataFirebird.destino)
			fmt.Scanf("h")
		}

	case 4:
	default:
		fmt.Println("Não foi escolhida uma opção válida. Sistema será encerrado.")
		os.Exit(1)
	}
}

func chooseOption(i *int) {

	fmt.Println("\n1 -\tBackup em Firebird\n2 - \tBackup em Postgres\n3 - \tRestore em Firebird\n4 - \tRestore em Postgres")
	fmt.Print("\nDigite a opção escolhida : ")
	fmt.Scanln(i)
	fmt.Println()

}

func getData(data *Data) {
	fmt.Println(data)

	in := bufio.NewReader(os.Stdin)

	if data.tipo == "FB" {

		fmt.Print("\nUsuario : ")
		line, _ := in.ReadString('\n')
		data.usuario = line

		fmt.Print("\nCaminho para o arquivo (ex : C:\\Spk Sistemas\\Bases\\BUSINESS.FDB) : ")
		line, err := in.ReadString('\n')
		if err != nil {
			panic("Falha ao obter dados referentes ao caminho para o arquivo.")
		}
		data.origem = line

		fmt.Print("\nCaminho onde será colocado o backup (ex : C:\\Spk Sistemas\\Backups\\BUSINESS.FDB) : ")
		line, err = in.ReadString('\n')
		if err != nil {
			panic("Falha ao obter dados referentes ao caminho onde será colocado o backup.")
		}
		data.destino = line

	} else {
		fmt.Print("\nUsuario : ")
		line, _ := in.ReadString('\n')
		data.usuario = line

		fmt.Print("\nNome do database no sistema : ")
		line, _ = in.ReadString('\n')
		data.database = line

		fmt.Print("\nCaminho onde será colocado o backup (ex : C:\\Backup\\rzbusiness.backup) : ")
		line, _ = in.ReadString('\n')
		data.destino = line

		fmt.Print("\nPorta : ")
		line, _ = in.ReadString('\n')
		data.porta = line
	}

	fixPath(data)

}

func msgScreen(msg string) {
	fmt.Println(strings.Repeat("#", 10), msg, strings.Repeat("#", 10))
}

func showDataFirebird(data *Data) {
	msgScreen("DADOS ATUAIS")
	var op string

	fmt.Printf("Usuario:\t%s\nNome Banco:\t%s\nOrigem Banco:\t%s\nDestino Banco:\t%s\n\n", data.usuario, data.database, data.origem, data.destino)
	fmt.Print("Você deseja alterar essas configurações ? (S/N) ")
	fmt.Scanln(&op)
	op = strings.ToUpper(op)
	if op == string('S') {
		getData(data)
	}
	fmt.Println()
}

func showDataPostgres(data *Data) {
	msgScreen("DADOS ATUAIS")
	var op string

	fmt.Printf("Usuario:\t%s\nNome Banco:\t%s\nDestino Banco:\t%s\n\n", data.usuario, data.database, data.destino)
	fmt.Print("Você deseja alterar essas configurações ? (S/N) ")
	fmt.Scanln(&op)
	op = strings.ToUpper(op)
	if op == string('S') {
		getData(data)
	}
	fmt.Println()
}

func fixPath(data *Data) {

	data.destino = strings.ReplaceAll(data.destino, "\n", "")
	data.origem = strings.ReplaceAll(data.destino, "\n", "")
	data.porta = strings.ReplaceAll(data.destino, "\n", "")
	data.database = strings.ReplaceAll(data.destino, "\n", "")
	data.usuario = strings.ReplaceAll(data.destino, "\n", "")
	data.tipo = strings.ReplaceAll(data.destino, "\n", "")
	data.destino = strings.ReplaceAll(data.destino, "\\", "\\\\")
	data.origem = strings.ReplaceAll(data.origem, "\\", "\\\\")
}
