package graph

func NumIslands(grid [][]byte) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 49 {
				result++
				grid = markIsland(grid, i, j)
			}
		}
	}
	return result
}

func markIsland(grid [][]byte, i, j int) [][]byte {
	q := [][]int{{i, j}}
	var current []int
	for len(q) > 0 {
		current = q[0]
		q = q[1:]

		if current[0] < 0 || current[0] >= len(grid) || current[1] < 0 || current[1] >= len(grid[current[0]]) || grid[current[0]][current[1]] != 49 {
			continue
		}

		grid[current[0]][current[1]] = 48
		q = append(q, []int{current[0] - 1, current[1]})
		q = append(q, []int{current[0] + 1, current[1]})
		q = append(q, []int{current[0], current[1] - 1})
		q = append(q, []int{current[0], current[1] + 1})
	}
	return grid
}
