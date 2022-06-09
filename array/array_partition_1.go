package main

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func arrayPairSum(nums []int) int {
	r := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i = i + 2 {
		r += min(nums[i-1], nums[i])
	}
	return r
}
