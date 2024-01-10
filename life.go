package main

import (
	"fmt"
	"time"
)

func main() {
	// setup_board
	//

	board := [] [3]byte {
		[3]byte{' ', ' ', ' '},
		[3]byte{' ', ' ', ' '},
		[3]byte{' ', ' ', ' '},
	}

	// create_first_pattern
	//

	board = [] [3]byte {
		[3]byte{' ', ' ', '*'},
		[3]byte{' ', '*', ' '},
		[3]byte{'*', ' ', ' '},
	}

	// print board
	//
	fmt.Printf("\n\n------\n")
    for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Printf("\n")
    }

	// ------
	time.Sleep(2 * time.Second)

	// mutate_board
	//

	// print board
	//
	fmt.Printf("\n\n------\n")
    for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Printf("%c", board[i][j])
		}
		fmt.Printf("\n")
    }
}
