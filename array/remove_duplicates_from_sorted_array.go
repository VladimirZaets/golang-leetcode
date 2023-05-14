package array

import "fmt"

func RemoveDuplicates(nums []int) int {
	prev := -101
	li := 0
	for i := 0; i < len(nums); i++ {
		if prev == nums[i] {
			nums[i] = -101
		} else {
			prev = nums[i]
			nums[li], nums[i] = nums[i], nums[li]
			li++
		}
	}
	return li
}

func RemoveDuplicates2(nums []int) int {
	l := len(nums)
	s := 1
	for i := 1; i < l; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		nums[s] = nums[i]
		s++
	}
	fmt.Println(nums)
	return s
}
