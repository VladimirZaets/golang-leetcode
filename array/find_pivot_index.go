package array

func FindPivotIndex(nums []int) int {
	sum := 0
	sumLeft := 0
	sumWithoutCurrent := 0
	result := -1
	for _, v := range nums {
		sum += v
	}
	for i := 0; i < len(nums); i++ {
		sumWithoutCurrent = sum - nums[i]
		if sumLeft*2 == sumWithoutCurrent {
			result = i
			break
		}
		sumLeft += nums[i]

	}
	return result
}
