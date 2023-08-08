package dynamic_programming

func CombinationSum4(nums []int, target int) int {
	return combinationSum4Memo(nums, target, map[int]int{})
}

func combinationSum4Memo(nums []int, target int, memo map[int]int) int {
	result := 0
	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target == 0 {
		result++
		return result
	}
	if target < 0 {
		return result
	}
	for _, v := range nums {
		result += combinationSum4Memo(nums, target-v, memo)
	}

	memo[target] = result
	return result
}
