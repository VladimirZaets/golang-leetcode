package main

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l

	left := nums[:l-k]
	list := nums[l-k:]
	list = append(list, left...)
	copy(nums, list)
}
