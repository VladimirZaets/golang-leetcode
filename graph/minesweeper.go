package graph

func UpdateBoard(board [][]byte, click []int) [][]byte {
	q := [][]int{click}
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	var current []int

	for len(q) > 0 {
		current = q[0]
		q = q[1:]
		nb := getNeighbors(board, current)
		mineQty := 0
		if board[current[0]][current[1]] == 'E' {
			for _, v := range nb {
				if board[v[0]][v[1]] == 'M' {
					mineQty++
				}
			}
			if mineQty > 0 {
				board[current[0]][current[1]] = byte(mineQty + '0')
			} else {
				board[current[0]][current[1]] = 'B'
				q = append(q, nb...)
			}
		}
	}

	return board
}

func getNeighbors(board [][]byte, current []int) [][]int {
	if (current[0] < 0 || current[0] >= len(board)) || (current[1] < 0 || current[1] >= len(board[current[0]])) {
		return [][]int{}
	}
	result := [][]int{}
	if current[0]-1 >= 0 && current[1]-1 >= 0 {
		result = append(result, []int{current[0] - 1, current[1] - 1})
	}
	if current[0]-1 >= 0 {
		result = append(result, []int{current[0] - 1, current[1]})
	}
	if current[0]-1 >= 0 && current[1]+1 < len(board[current[0]]) {
		result = append(result, []int{current[0] - 1, current[1] + 1})
	}
	if current[1]-1 >= 0 {
		result = append(result, []int{current[0], current[1] - 1})
	}
	if current[1]+1 < len(board[current[0]]) {
		result = append(result, []int{current[0], current[1] + 1})
	}
	if current[0]+1 < len(board) && current[1]-1 >= 0 {
		result = append(result, []int{current[0] + 1, current[1] - 1})
	}
	if current[0]+1 < len(board) {
		result = append(result, []int{current[0] + 1, current[1]})
	}
	if current[0]+1 < len(board) && current[1]+1 < len(board[current[0]]) {
		result = append(result, []int{current[0] + 1, current[1] + 1})
	}
	return result
}
