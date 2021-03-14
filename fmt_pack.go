package main

import (
	"fmt"
)

func main(){

	x := "Hello\nGood Morning!"
	y := `How are you?`
	z := fmt.Sprint(x, " ", y);
	fmt.Println(z);
}

/*
	#1: Print -> Standard out
			func Print(a ...interface{}) (n int, err error)
			func Printl - (new line )
			func Printf(format string, a ... interface{}) (n int, err error)
				- Format verbs (%v %T)
	#2: Print -> string (can be used like variable)
			func Sprint(a ...interface{}) string
			func Sprintf(format string, a ...interface{}) string
			func Sprintln(a ...interface{}) string
	#3: Print -> file, writer interface, e.g. (archives or response from server)
			func Fprint(w io.Writer, a ...interface{}) (n int, err error)
			func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
			func Fprintln(w io.Writer, a ...interface{})
*/