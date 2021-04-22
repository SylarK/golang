package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main(){

	path := "C:\\Programdata\\SPK Sistemas\\alias.txt"

	retorno, err := getText(path)
	if err != nil{
		fmt.Println("Erro ao recuperar informações : ", err)
	}

	for i, _ := range retorno{
		retorno[i] = retorno[i][strings.Index(retorno[i], "=") + 1:]
	}

	fmt.Println(retorno)

}

func getText(path string) ([]string , error){

	arquivo, err := os.Open(path)
	if err != nil{
		fmt.Println("Erro ao abrir o caminho especificado.\n", err)
	}
	defer arquivo.Close()

	var lines []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

/*
func main() {
	cmd := exec.Command("sleep", "1")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}
*/