package dynamic_programming

import (
	"fmt"
	"strconv"
)

func SolveSudoku(board [][]byte) {
	FindValid(board, 0, 0)

	for _, row := range board {
		r := ""
		for _, cell := range row {
			r += fmt.Sprintf("%s ", string(cell))
		}
		fmt.Println(r)
	}
}

func FindValid(board [][]byte, row int, col int) bool {
	if col == 9 {
		row += 1
		col = 0
	}

	if row == 9 {
		return true
	}

	if board[row][col] != '.' {
		return FindValid(board, row, col+1)
	}

	for i := 1; i < 10; i++ {
		v := []byte(strconv.Itoa(i))[0]
		board[row][col] = v
		isCellValid := validatePoint(board, row, col)
		if isCellValid {
			if FindValid(board, row, col+1) {
				return true
			}
		}
		board[row][col] = '.'
	}

	return false
}

func validatePoint(board [][]byte, row int, col int) bool {
	indexes := [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}
	v := board[row][col]
	for i := 0; i < 9; i++ {
		if i != row && board[i][col] == v {
			return false
		}
		if i != col && board[row][i] == v {
			return false
		}
	}
	getSet := func(v int) int {
		if v < 3 {
			return 0
		} else if v >= 3 && v < 6 {
			return 3
		}

		return 6
	}
	rowSetStart := getSet(row)
	coluntSetStart := getSet(col)
	for _, item := range indexes {
		if item[0]+rowSetStart != row && item[1]+coluntSetStart != col && board[item[0]+rowSetStart][item[1]+coluntSetStart] == v {
			return false
		}
	}

	return true
}
