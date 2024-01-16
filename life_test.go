package main

import (
    "testing"
)

// returns an 0 or 1 depending on whether this adds to the
// count of alive neighbours
func Test_cellAliveIntd(t *testing.T) {
	var board lifeBoard;

	switch cellAliveInt(board, -1, -1) {
	case 0:
	default: t.Errorf("out of bounds call: cellAliveInt(board, -1, -1) should return 0")
	}

	board[0][0] = ' '
	switch cellAliveInt(board, 0, 0) {
	case 0:
	default: t.Errorf(": cellAliveInt(...) should return 0, when cell is inbounds and dead")
	}

	board[0][0] = '@'
	switch cellAliveInt(board, 0, 0) {
	case 1:
	default: t.Errorf(": cellAliveInt(...) should return 1, when cell is inbounds and alive")
	}
}

func TestisCellOnBoard(t *testing.T) {
   var tests = []struct {
        x, y int
        want bool
    }{
        {0, 0, true},
        {7, 7, true},
        {-1, 0, false},
        {0, -1, false},
        {8, 0, false},
        {0, 8, false},
    }

    for _, tt := range tests {
		if (isCellOnBoard(tt.x, tt.y)) != tt.want {
			t.Errorf("isCellOnBoard(%d, %d) should be %t", tt.x, tt.y, tt.want)
		}
	}
}

func Test_isCellAlive(t *testing.T) {
	var board lifeBoard;

	board[0][0] = '@'
    if isCellAlive(board, 0, 0) != true {
		t.Errorf("isCellAlive(...) should be true for cell with '@'")
	}

	board[0][0] = ' '
    if isCellAlive(board, 0, 0) != false {
		t.Errorf("isCellAlive(...) should be false for cell with ' '")
	}

	// NOTE: not sure how to test
	// board[0][0] = '!'
    // if isCellAlive(board, 0, 0) != 0 {
	// 	t.Errorf("isCellAlive(...) should explode for cell with '!'")
	// }
}
