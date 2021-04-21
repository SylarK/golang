//writing files

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	d1 := []byte("Hello\nGophers\n")
	err := ioutil.WriteFile("./sample_arquivo.txt", d1, 0604)
	checkError(err)
	//
	f, err := os.Create("./tmp.txt")
	checkError(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	checkError(err)
	fmt.Printf("Wrote %d bytes\n", n4)

	w.Flush()

}
