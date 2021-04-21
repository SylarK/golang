package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type erro error

func checkError(err erro, i int) {
	if err != nil {
		fmt.Println("Erro : #", i)
		panic(err)
	}
}

func main() {

	var path = "C:/ProgramData/go/sample.txt"

	data, err := ioutil.ReadFile(path)
	checkError(err, 1)

	fmt.Println(string(data))

	f, err := os.Open(path)
	checkError(err, 2)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	checkError(err, 3)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(5, 0)
	checkError(err, 3)
	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	checkError(err, 4)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n: ", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	checkError(err, 5)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	checkError(err, 6)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	checkError(err, 7)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	checkError(err, 8)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()

}

/*
Atenção.

1. Read reads up to len(b) bytes from the File.
It returns the number of bytes read and any error encountered. At end of file, Read returns 0, io.EOF.
*/
