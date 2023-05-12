package array

func FindMaxConsecutiveOnes(nums []int) int {
	max := 0
	current := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			current++
			if current > max {
				max = current
			}
		} else {
			current = 0
		}
	}

	return max
}
