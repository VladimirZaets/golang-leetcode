package main

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			prev := result[i-1]
			var leftOperand, rightOperand int
			if j-1 < 0 {
				leftOperand = 0
			} else {
				leftOperand = prev[j-1]
			}

			if j > len(prev)-1 {
				rightOperand = 0
			} else {
				rightOperand = prev[j]
			}
			result[i][j] = leftOperand + rightOperand
		}
	}
	return result
}
