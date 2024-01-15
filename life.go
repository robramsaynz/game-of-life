// An implements Conway's Game of Life:
//
// - https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#Rules
// - https://pi.math.cornell.edu/~lipa/mec/lesson6.html

package main

// NOTE 1
//
// For [x][y] coordinates [0][0] == bottom left, and
// [2][2] == top-right.
//
// Generally this means that y works how you expect,
// but y inverted for things like printing, for which
// the highest line is printed first.

import (
	"fmt"
	"time"
	"log"
)

const boardSizeX = 8
const boardSizeY = 8

type lifeBoard [boardSizeX][boardSizeY]byte;

func main() {
	board := setup_board()
	printBoard(board)
	time.Sleep(time.Second)

	for i := 0 ; i < 10 ; i++ {
		board = processBoard(board)
		printBoard(board)
		time.Sleep(time.Second)
	}
}

func setup_board() lifeBoard {
	// This is inverted vertically (see NOTE 1)
	board := [boardSizeX][boardSizeY]byte {
		[boardSizeY]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		[boardSizeY]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		[boardSizeY]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		[boardSizeY]byte{' ', ' ', '@', '@', '@', ' ', ' ', ' '},
		[boardSizeY]byte{' ', '@', '@', '@', ' ', ' ', ' ', ' '},
		[boardSizeY]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		[boardSizeY]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		[boardSizeY]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
	}

	return board
}

func clearScreen() {
	fmt.Printf("\033[H\033[2J")
}

func printBoard(board lifeBoard) {
	clearScreen()

	fmt.Printf("\n-------------------\n")

	// NOTE: x--, and y++ (see NOTE 1)
    for x := boardSizeX-1; x >= 0; x-- {
		fmt.Printf("| ")
		for y := 0; y <= boardSizeY-1; y++ {
			fmt.Printf(" %c", board[x][y])
		}
		fmt.Printf(" |\n")
    }

	fmt.Printf("-------------------\n")
}


func processBoard(oldBoard lifeBoard) lifeBoard {
	var newBoard lifeBoard

	// NOTE: x--, and y++ (see NOTE 1)
    for x :=  boardSizeX-1; x >= 0; x-- {
		for y := 0; y <= boardSizeY-1; y++ {
			newBoard[x][y] = processCell(oldBoard, x, y)
		}
    }

	return newBoard
}

func processCell(board lifeBoard, x int, y int) byte {
	totalLiveNeighbours :=
		cellAliveInt(board, x, y+1) +   // top-cell
		cellAliveInt(board, x, y-1) +   // bottom-cell
		cellAliveInt(board, x-1, y) +   // left-cell
		cellAliveInt(board, x+1, y) +   // right-cell
		cellAliveInt(board, x-1, y+1) + // top-left-cell
		cellAliveInt(board, x+1, y+1) + // top-right-cell
		cellAliveInt(board, x+1, y-1) + // bottom-right-cell
		cellAliveInt(board, x-1, y-1)   // bottom-left-cell

	var cellLives bool

	isThisCellAlive := isCellAlive(board, x, y)

	switch isThisCellAlive {
	case true:
		switch {
		// Any live cell with fewer than two live neighbours dies, as if by underpopulation
		case totalLiveNeighbours < 2: cellLives = false
		// Any live cell with more than three live neighbours dies, as if by overpopulation.
		case totalLiveNeighbours > 3: cellLives = false
		// Any live cell with two or three live neighbours lives on to the next generation.
		default: cellLives = true
		}
	case false:
		switch totalLiveNeighbours {
		// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
		case 3: cellLives = true
		default: cellLives = false
		}
	}

	switch cellLives {
	case false: return ' '
	case true:  return '@'
    default: log.Fatal("invalid switch state") ; return '?'
	}
}

// returns an 0 or 1 depending on whether this adds to the
// count of alive neighbours
func cellAliveInt(board lifeBoard, x int, y int) int {
	switch {
	case ! isCellOnBoard(x, y): return 0
	case ! isCellAlive(board, x, y): return 0
	default: return 1
	}
}

func isCellOnBoard(x, y int) bool {
	switch {
	case x < 0: return false
	case x >= boardSizeX-1: return false
	case y < 0: return false
	case y > boardSizeY-1: return false
	default: 	return true
	}
}

func isCellAlive(board lifeBoard, x int, y int) bool{
    switch board[x][y] {
    case '@': return true
    case ' ': return false
    default: log.Fatal("invalid switch arg: ", board[x][y]) ; return false
    }
}
