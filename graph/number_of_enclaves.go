package graph

func NumEnclaves(grid [][]int) int {
	result := 0
	cellsQty := func(grid [][]int, current []int) int {
		q := [][]int{current}
		isCorrect := true
		result := 0
		for len(q) > 0 {
			current = q[0]
			q = q[1:]
			if grid[current[0]][current[1]] != 1 {
				continue
			}

			if current[0] == 0 || current[0] == len(grid)-1 || current[1] == 0 || current[1] == len(grid[current[0]])-1 || grid[current[0]][current[1]] != 1 {
				isCorrect = false
				continue
			}
			grid[current[0]][current[1]] = 0
			result++

			q = append(q, []int{current[0] - 1, current[1]})
			q = append(q, []int{current[0] + 1, current[1]})
			q = append(q, []int{current[0], current[1] - 1})
			q = append(q, []int{current[0], current[1] + 1})
		}
		if isCorrect {
			return result
		}
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				result += cellsQty(grid, []int{i, j})
			}
		}
	}
	return result
}
