package graph

func ClosedIsland(grid [][]int) int {
	result := 0
	isIsland := func(grid [][]int, current []int) bool {
		q := [][]int{current}
		result := true
		for len(q) > 0 {
			current = q[0]
			q = q[1:]

			if grid[current[0]][current[1]] == 1 {
				continue
			}

			if current[0] == 0 || current[0] >= len(grid)-1 || current[1] == 0 || current[1] >= len(grid[current[0]])-1 {
				result = false
				continue
			}

			grid[current[0]][current[1]] = 1
			q = append(q, []int{current[0] - 1, current[1]})
			q = append(q, []int{current[0] + 1, current[1]})
			q = append(q, []int{current[0], current[1] - 1})
			q = append(q, []int{current[0], current[1] + 1})
		}
		return result
	}

	for i, row := range grid {
		for j, val := range row {
			if val == 0 && i != 0 && i != len(grid)-1 && j != 0 && j != len(row)-1 {
				r := isIsland(grid, []int{i, j})
				if r {
					result++
				}
			}
		}
	}

	return result
}
