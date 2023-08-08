package dynamic_programming

func IsValidSudoku(board [][]byte) bool {
	for i, row := range board {
		for j, cell := range row {
			if cell == '.' {
				continue
			}
			if !validate(board, i, j) {
				return false
			}
		}
	}
	return true
}

func validate(board [][]byte, row int, col int) bool {
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
