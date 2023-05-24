package array

import "sort"

func HeightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)
	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sorted[i] {
			count++
		}
	}
	return count
}
