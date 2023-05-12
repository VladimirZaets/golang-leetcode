package array

func SpiralMatrix(matrix [][]int) []int {
	const (
		LEFT = iota
		RIGHT
		UP
		DOWN
	)
	iterations := len(matrix[0]) * len(matrix)
	direction := RIGHT
	col := 0
	row := 0
	passedRowsMin := -1
	passedColsMin := -1
	passedColsMax := len(matrix[0]) - 1
	passedRowsMax := len(matrix) - 1
	result := make([]int, 0)
	var down, right, left, up func()

	right = func() {
		if col+1 <= passedColsMax {
			col = col + 1
		} else {
			direction = DOWN
			if passedRowsMin-passedRowsMax == 0 {
				return
			}
			passedRowsMin = passedRowsMin + 1

			down()
		}
	}
	left = func() {
		if col-1 > passedColsMin {
			col = col - 1
		} else {
			direction = UP
			if passedRowsMin-passedRowsMax == 0 {
				return
			}
			passedRowsMax = passedRowsMax - 1
			up()
		}
	}
	down = func() {
		if row+1 <= passedRowsMax {
			row = row + 1
		} else {
			direction = LEFT
			if passedColsMax-passedColsMin == 0 {
				return
			}
			passedColsMax = passedColsMax - 1
			left()
		}
	}
	up = func() {
		if row-1 > passedRowsMin {
			row = row - 1
		} else {
			direction = RIGHT
			if passedColsMax-passedColsMin == 0 {
				return
			}
			passedColsMin = passedColsMin + 1
			right()
		}
	}

	for i := 0; i < iterations; i++ {

		result = append(result, matrix[row][col])

		switch direction {
		case LEFT:
			left()

		case RIGHT:
			right()

		case DOWN:
			down()

		case UP:
			up()

		}

	}

	return result
}
