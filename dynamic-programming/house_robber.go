package dynamic_programming

func Rob(nums []int) int {
	return rob(nums, 0, map[int]int{})
}

func rob(nums []int, start int, memo map[int]int) int {
	if _, ok := memo[start]; ok {
		return memo[start]
	}
	if start >= len(nums) {
		return 0
	}

	first := rob(nums, start+1, memo)
	second := nums[start] + rob(nums, start+2, memo)

	if first > second {
		memo[start] = first
		return memo[start]
	}
	memo[start] = second
	return memo[start]
}
