package array

func MinSubArrayLen(target int, nums []int) int {
	sum := 0
	min := len(nums) + 1
	j := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		for sum >= target {
			sum -= nums[j]
			c := (i + 1) - j
			if min > c {
				min = (i + 1) - j
			}
			j++
		}
	}

	if min > len(nums) {
		return 0
	}
	return min
}
