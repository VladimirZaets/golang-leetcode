package dynamic_programming

func PascalsTriangleGenerate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}

	if numRows == 1 {
		return result
	}
	result[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		item := make([]int, i+1)
		item[0] = 1
		item[i] = 1
		for j := 1; j < i; j++ {
			item[j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i] = item
	}
	return result
}
