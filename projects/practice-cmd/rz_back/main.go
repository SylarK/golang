package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
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
	host     string
}

func main() {

	var option int
	var err error
	caminho := "C:\\Programdata\\SPK Sistemas\\aliasSPK.ini"
	retorno, falha := getText(caminho)
	if falha != nil {
		log.Fatalf("Error ao buscar C:\\Programdata\\SPK Sistemas\\aliasSPK.ini\nVerifique se o arquivo está criado.\n\n%v", falha.Error())
	}

	for i, _ := range retorno {
		retorno[i] = retorno[i][strings.Index(retorno[i], "=")+1:]
	}

	dataFirebird := Data{
		"FB", retorno[5], "spk159288", "BUSINESS", retorno[3], "C:\\Spk Sistemas\\BKPRZ_botFB.FBK", retorno[7], retorno[4],
	}

	dataPostgres := Data{
		"PG", retorno[5], "-", "BUSINESS", retorno[3], "C:\\Spk Sistemas\\BKPRZ_botPG.backup", retorno[7], retorno[4],
	}

	//showData(&data)
	msgScreen("RAZEM")
	chooseOption(&option)

	switch option {
	case 1:
		clearScreen()
		showDataFirebird(&dataFirebird)
		command := exec.Command("gbak.exe", "-b", "-user", dataFirebird.usuario, "-pas", dataFirebird.senha, dataFirebird.origem, dataFirebird.destino)
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
			fmt.Printf("\nComando executado com sucesso. Aguarde o backup ficar pronto.\nCaminho do arquivo : %s\n\nPressione 'Enter' para sair...", strings.ReplaceAll(dataFirebird.destino, "\\\\", "\\"))
			fmt.Scanf("h")
		}
	case 2:
		clearScreen()
		showDataPostgres(&dataPostgres)
		command := exec.Command("pg_dump.exe", "--host", "localhost", "--port", "5432", "--username", dataPostgres.usuario, "--format", "tar", "--file", dataPostgres.destino, dataPostgres.database)
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
			fmt.Printf("\nComando executado com sucesso. Aguarde o backup ficar pronto.\nCaminho do arquivo : %s\n\nPressione 'Enter' para sair...", strings.ReplaceAll(dataPostgres.destino, "\\\\", "\\"))
			fmt.Scanf("h")
		}
	case 3:
		clearScreen()
		fmt.Print("Caminho origem do arquivo a ser restaurado (ex : C:\\Spk Sistemas\\Backups\\BACKUP.FBK) : ")
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')
		dataFirebird.origem = line[0 : len(line)-2]
		fmt.Print("Caminho destino do arquivo a ser restaurado (ex : C:\\Spk Sistemas\\Bases\\RZBUSINESS_RESTAURADO.FDB) : ")
		line, _ = in.ReadString('\n')
		dataFirebird.destino = line[0 : len(line)-2]
		fixPath(&dataFirebird)

		command := exec.Command("gbak.exe", "-r", "-user", dataFirebird.usuario, "-password", dataFirebird.senha, dataFirebird.origem, dataFirebird.destino)
		fmt.Println("\nAguarde, executando comando...")
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
			fmt.Printf("\nComando executado com sucesso. Aguarde o seu restore ficar pronto.\nCaminho do arquivo : %s\n\nPressione 'Enter' para sair...", strings.ReplaceAll(dataFirebird.destino, "\\\\", "\\"))
			fmt.Scanf("h")
		}

	case 4:
		clearScreen()
		fmt.Print("* Atenção : Lembre-se que para se utilizar esse restore é necessário que exista um banco de dados já cadastrado no postgres para restaurar as informações diretamente para o banco.\n\n")
		fmt.Print("Caminho origem do arquivo a ser restaurado (ex : C:\\Spk Sistemas\\Backups\\BACKUP.backup) : ")
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')
		dataPostgres.origem = line[0 : len(line)-2]

		fmt.Print("\nPorta : ")
		line, _ = in.ReadString('\n')
		dataPostgres.porta = line[0 : len(line)-2]

		fmt.Print("\nUsuário : ")
		line, _ = in.ReadString('\n')
		dataPostgres.usuario = line[0 : len(line)-2]

		fmt.Print("\nNome do Banco de dados já cadastrado : ")
		line, _ = in.ReadString('\n')
		dataPostgres.database = line[0 : len(line)-2]

		command := exec.Command("pg_restore.exe", "--host", "localhost", "--port", dataPostgres.porta, "--username", dataPostgres.usuario, "--dbname", dataPostgres.database, dataPostgres.origem)
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
			fmt.Printf("\nComando executado com sucesso. A restauração foi iniciada.\n\nPressione 'Enter' para sair...")
			fmt.Scanf("h")
		}

	default:
		clearScreen()
		fmt.Println("A opção escolhida não é válida. O sistema será encerrado.")
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

	msgScreen("ALTERAÇÃO")

	in := bufio.NewReader(os.Stdin)

	if data.tipo == "FB" {

		fmt.Print("\nUsuario : ")
		line, _ := in.ReadString('\n')
		data.usuario = line[0 : len(line)-2]

		fmt.Print("\nPassword : ")
		line, _ = in.ReadString('\n')
		data.senha = line[0 : len(line)-2]

		fmt.Print("\nCaminho para o arquivo (ex : C:\\Spk Sistemas\\Bases\\BUSINESS.FDB) : ")
		line, _ = in.ReadString('\n')
		data.origem = line[0 : len(line)-2]

		fmt.Print("\nCaminho onde será colocado o backup (ex : C:\\Spk Sistemas\\Backups\\BACKUP.FBK) : ")
		line, _ = in.ReadString('\n')
		data.destino = line[0 : len(line)-2]

		fmt.Print("\nPorta : ")
		line, _ = in.ReadString('\n')
		data.porta = line[0 : len(line)-2]

	} else {
		fmt.Print("\nUsuario : ")
		line, _ := in.ReadString('\n')
		data.usuario = line[0 : len(line)-2]

		fmt.Print("\nNome do database no sistema : ")
		line, _ = in.ReadString('\n')
		data.database = line[0 : len(line)-2]

		fmt.Print("\nCaminho onde será colocado o backup (ex : C:\\Backup\\rzbusiness.backup) : ")
		line, _ = in.ReadString('\n')
		data.destino = line[0 : len(line)-2]
		data.destino = strings.ReplaceAll(data.destino, "\\", "\\\\")

		fmt.Print("\nPorta : ")
		line, _ = in.ReadString('\n')
		data.porta = line[0 : len(line)-2]
	}

	fixPath(data)

}

func msgScreen(msg string) {
	fmt.Println()
	fmt.Println(strings.Repeat("#", 10), msg, strings.Repeat("#", 10))
}

func showDataFirebird(data *Data) {
	msgScreen("DADOS ATUAIS")
	var op string

	fmt.Printf("Usuario:\t%s\nNome Banco:\t%s\nOrigem Banco:\t%s\nDestino Banco:\t%s\nPorta:\t\t%s\n\n", data.usuario, data.database, data.origem, data.destino, data.porta)
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

	fmt.Printf("Usuario:\t%s\nNome Banco:\t%s\nDestino Banco:\t%s\nPorta:\t\t%s\n", data.usuario, data.database, data.destino, data.porta)
	fmt.Print("Você deseja alterar essas configurações ? (S/N) ")
	fmt.Scanln(&op)
	op = strings.ToUpper(op)
	if op == string('S') {
		getData(data)
	}
	fmt.Println()
}

func fixPath(data *Data) {
	data.destino = strings.ReplaceAll(data.destino, "\\", "\\\\")
	data.origem = strings.ReplaceAll(data.origem, "\\", "\\\\")
}

func clearScreen() {
	command := exec.Command("cmd", "/c", "cls")
	command.Stdout = os.Stdout
	command.Run()
}

func getText(path string) ([]string, error) {

	arquivo, _ := os.Open(path)
	/* if err != nil{
		fmt.Println("Erro ao abrir o caminho especificado.\n", err)
	} */
	defer arquivo.Close()

	var lines []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()

}

//Lucas Amado - @amadodev
