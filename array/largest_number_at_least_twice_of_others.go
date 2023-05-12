package array

func DominantIndex(nums []int) int {
	maxNumber := -1
	maxIndex := -1
	for i, v := range nums {
		if v >= maxNumber*2 {
			maxNumber = v
			maxIndex = i
		} else if v > maxNumber {
			maxNumber = v
			maxIndex = -1
		} else if v > maxNumber/2 {
			maxIndex = -1
		}
	}
	return maxIndex
}
