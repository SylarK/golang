// Writing data to a TCP connection

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
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil{
		log.Fatal("Erro ao iniciar tcp server : ", err)
	}

	defer listener.Close()

	log.Println("Listening on " + HOST + ":" + PORT)

	for{
		conn, err := listener.Accept()
		if err != nil{
			log.Fatal("Error acception: ", err.Error())
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn){
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil{
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Print("Message Received : ", string(message))
	conn.Write([]byte(message + "\n"))
	conn.Close()
}

/*
2021/05/02 17:59:39 Listening on localhost:8080

<-- "Hello TCP Server. [again]" | ncat localhost 8080
--> "Hello TCP Server. [again]"

--> Message Received : "Hello TCP Server. [again]"
*/