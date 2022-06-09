package main

func getRow(rowIndex int) []int {
	prev := []int{1}
	var current []int

	for i := 1; i <= rowIndex; i++ {
		for j := 0; j <= i; j++ {
			if j >= len(prev) {
				current = append(current, prev[j-1])
			} else if j-1 >= 0 {
				current = append(current, prev[j]+prev[j-1])
			} else {
				current = append(current, prev[j])
			}
		}
		prev = current
		current = make([]int, 0)
	}

	return prev
}
