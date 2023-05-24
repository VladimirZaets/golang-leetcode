package array

func ThirdMaximumNumber(nums []int) int {
	max := nums[0]
	emptyInt := -999999999999999999
	second := emptyInt
	third := emptyInt

	for i := 1; i < len(nums); i++ {
		if max == nums[i] || second == nums[i] || third == nums[i] {
			continue
		}
		if nums[i] > max {
			third = second
			second = max
			max = nums[i]
		} else if second == emptyInt || nums[i] > second {
			third = second
			second = nums[i]
		} else if third == emptyInt || nums[i] > third {
			third = nums[i]
		}
	}
	if third == emptyInt {
		return max
	}

	return third
}
