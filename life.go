package main

import (
	"fmt"
	"time"
)

func main() {
	board := setup_board()

	printBoard(board)

	// ------
	time.Sleep(2 * time.Second)

	// mutate_board
	//
	board[0][0] = '*'
	printBoard(board)
}

func setup_board() [3][3]byte  {
	board := [3] [3]byte {
		[3]byte{' ', ' ', '*'},
		[3]byte{' ', '*', ' '},
		[3]byte{'*', ' ', ' '},
	}

	return board
}

func clearScreen() {
	fmt.Printf("\033[H\033[2J")
}

func printBoard(board [3][3]byte) {
	clearScreen()

	fmt.Printf("\n------\n")

    for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Printf("\n")
    }

	fmt.Printf("------\n")
}
