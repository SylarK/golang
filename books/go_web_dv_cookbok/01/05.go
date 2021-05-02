//Reading data from a TCP connection

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const(
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main(){
	listener, err := net.Listen(TYPE, HOST + ":" + PORT)
	if err != nil{
		log.Fatal("Erro ao iniciar o servidor TCP : ", err)
	}
	defer listener.Close()

	log.Println("Listening " + HOST + ":" + PORT)

	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Fatal("Error accepting: ", err.Error())
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn){
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil{
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Print("Mensagem recebida do cliente: ", string(message))
	conn.Close()
}

/*
2021/05/02 17:45:13 Listening localhost:8080

<-- echo -n "Hello TCP Server" | ncat localhost 8080

--> Mensagem recebida do cliente: -n "Hello TCP Server"
*/