/*
	Channels provide a way for two goroutines to communicate 
	with one another and synchronize their execution.

	Channel type is represented with the keyword "chan" followed
	by the type of the things that are passed on the channel.
	The operator "<-" is used to send and receibe messages on
	the channel.
*/

package main

import (
	"fmt"
	"time"
)

//func pinger(c chan string)
func pinger(c chan<- string){ //c can only be sent to
	for i:= 0; ;i++{
		c <- "ping"
	}
}

func ponger(c chan string){
	for i:=0; ;i++{
		c <- "pong"
	}
}

func printer(c <-chan string){
	for{
		msg := <- c
		fmt.Println(msg)
		//fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func main(){
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}