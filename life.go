package main

import (
	"fmt"
	"time"
	"log"
)

type lifeBoard [3][3]byte;

// func cellIsLive(cell byte) {
func main() {
	board := setup_board()
	printBoard(board)

	time.Sleep(2 * time.Second)

	board = mutateBoard(board)
	printBoard(board)
}

func setup_board() lifeBoard {
	board := [3][3]byte {
		[3]byte{'.', '.', '@'},
		[3]byte{'.', '@', '.'},
		[3]byte{'@', '.', '.'},
	}

	return board
}

func clearScreen() {
	fmt.Printf("\033[H\033[2J")
}

func printBoard(board lifeBoard) {
	clearScreen()

	fmt.Printf("\n------\n")

    for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Printf(" %c", board[i][j])
		}
		fmt.Printf("\n")
    }

	fmt.Printf("------\n")
}


func mutateBoard(board lifeBoard) lifeBoard {
    for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			switch board[i][j] {
			case '.': board[i][j] = '@'
			case '@': board[i][j] = '.'
			default: log.Fatal("Invalid char in cell")
			}
		}
    }

	return board
}
