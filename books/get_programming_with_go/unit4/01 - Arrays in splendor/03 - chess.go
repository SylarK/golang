// array of arrays

package main

import "fmt"

func main() {
	var board [8][8]string
	board[0][0] = "r"
	board[0][7] = "r"
	for column := range board[1] {
		board[1][column] = "p"
	}
	for column := range board[6] {
		board[6][column] = "p"
	}
	fmt.Print(board)
}

// sudoku

func sudoku() {
	var grid [9][9]int
	fmt.Print(grid)
}
