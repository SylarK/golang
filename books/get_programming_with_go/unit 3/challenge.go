/*
Write a program that displays temperature conversion
tables. The tables should use equals signs (=) and pipes
(|) to draw lines
The program should draw two tables. The first table has
two columns, with C in the first column and F in the
second column.
After completing one table, implement a second table with the columns reversed, converting from F to C.
*/

package main

import "fmt"

type celsius float64
type fahrenheit float64

type getRowFn func(row int) (string, string)

// C -> F
func (c celsius) f() fahrenheit {
	return fahrenheit((c * 9 / 5) + 32)
}

// F -> C
func (f fahrenheit) c() celsius {
	return celsius((f - 32) * 5 / 9)
}

//
const (
	line      = "======================="
	format    = "%.1f"
	rowFormat = "| %8s | %8s |\n"
)

func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}
func ctof(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.f()
	cell1 := fmt.Sprintf(format, c)
	cell2 := fmt.Sprintf(format, f)
	return cell1, cell2
}
func ftoc(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.c()
	cell1 := fmt.Sprintf(format, f)
	cell2 := fmt.Sprintf(format, c)
	return cell1, cell2
}
func main() {
	/* var c celsius = 10.0
	var f fahrenheit = 150
	fmt.Println("Testing...")
	fmt.Printf("C -> F : %.2f\n", c.f())
	fmt.Printf("F -> C : %.2f", f.c()) */
	drawTable("ºC", "ºF", 5, ctof) //celsius to f
	fmt.Println()
	drawTable("ºF", "ºC", 5, ftoc) // f to celsius
}
