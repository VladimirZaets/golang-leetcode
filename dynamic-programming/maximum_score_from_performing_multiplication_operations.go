package dynamic_programming

import (
	"fmt"
)

func MaximumScore(nums []int, multipliers []int) int {
	return maximumScoreRec(nums, multipliers, 0, len(nums)-1, 0, 0, -1, map[string]int{})
}

func maximumScoreRec(nums []int, multipliers []int, left int, right, i int, sum int, depth int, memo map[string]int) int {
	lr := fmt.Sprintf("%d:%d", left, right)
	if _, ok := memo[lr]; ok {
		return memo[lr]
	}
	if i == len(multipliers) {
		return 0
	}

	if i == len(multipliers)-1 {
		if nums[left]*multipliers[i] > nums[right]*multipliers[i] {
			return nums[left] * multipliers[i]
		}
		return nums[right] * multipliers[i]
	}

	if left == len(nums)-1 {
		return multipliers[i] * nums[left]
	}

	if right == 0 {
		return multipliers[i] * nums[right]
	}

	s1 := multipliers[i]*nums[left] + maximumScoreRec(nums, multipliers, left+1, right, i+1, 0, depth+1, memo)
	s2 := multipliers[i]*nums[right] + maximumScoreRec(nums, multipliers, left, right-1, i+1, 0, depth+1, memo)

	if s1 > s2 {
		sum += s1
	} else {
		sum += s2
	}

	memo[lr] = sum

	return memo[lr]
}
