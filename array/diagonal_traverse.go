package main

func findDiagonalOrder(mat [][]int) []int {
	direction := 1
	col := 0
	row := 0
	totalRows := len(mat) - 1
	totalCols := len(mat[0]) - 1
	result := make([]int, 0)

	for i := 0; i < ((totalRows + 1) * (totalCols + 1)); i++ {
		result = append(result, mat[row][col])
		if direction == 1 {
			if row-1 >= 0 && col+1 <= totalCols {
				row = row - 1
				col = col + 1
				continue
			}

			if row-1 < 0 && col+1 <= totalCols {
				direction = 0
				col = col + 1
				continue
			}

			if col+1 > totalCols {
				direction = 0
				row = row + 1
				continue
			}
		} else {
			if row+1 <= totalRows && col-1 >= 0 {
				row = row + 1
				col = col - 1
				continue
			}

			if col-1 < 0 && row+1 <= totalRows {
				direction = 1
				row = row + 1
				continue
			}

			if row+1 > totalRows {
				direction = 1
				col = col + 1
				continue
			}

		}
	}
	return result
}
