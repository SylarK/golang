package main

import (
	"bufio"
	"fmt"
	"os"
)

type Variaveis struct{
	variavel1 string
	variavel2 string
	variavel3 string
}

func main(){

	variaveis := Variaveis{}

	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	variavel := line
	//variavel = variavel[0:len(variavel) - 2]
	variaveis.variavel1 = variavel;

	line, _ = in.ReadString('\n')
	variavel = line
	//variavel = variavel[0:len(variavel) - 2]
	variaveis.variavel2 = variavel;

	line, _ = in.ReadString('\n')
	variavel = line
	//variavel = variavel[0:len(variavel) - 2]
	variaveis.variavel3 = variavel;

	fmt.Println(variaveis)
}