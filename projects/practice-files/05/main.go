package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){

	path := "C:\\ProgramData\\SPK Sistemas\\alias.txt"
	retorno, err := getText(path)
	if err != nil {
		fmt.Println("Erro ao gerar retorno : ", err)
	}
	fmt.Println(retorno)

	var conteudo []string
	conteudo = append(conteudo, "123")
	conteudo = append(conteudo, "456")
	conteudo = append(conteudo, "789")

	err = writeText(conteudo, "foo.txt")
	if err != nil{
		log.Fatalf("Erro : ", err)
	}

	for indice, linha := range conteudo{
		fmt.Println(indice, linha)
	}
}

func getText(pathText string) ([]string, error){
	//Abrindo o arquivo
	arquivo, err := os.Open(pathText)
	if err != nil{
		fmt.Println("Caminho não encontrado.")
	}
	defer arquivo.Close()

	//Lendo o arquivo
	var lines []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func writeText(lines []string, pathText string) error{
	//Criando o arquivo de texto
	arquivo, err := os.Create(pathText)
	if err != nil{
		return err
	}
	defer arquivo.Close()
	
	//Escritor responsável por cada linha
	write := bufio.NewWriter(arquivo)
	for _, line := range lines{
		fmt.Fprintln(write, line)
	}

	return write.Flush()

} 