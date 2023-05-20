package array

import "fmt"

func SortArrayByParity(nums []int) []int {
	s := 0
	e := len(nums) - 1

	for s < e {
		if nums[s]%2 > nums[e]%2 {
			nums[s], nums[e] = nums[e], nums[s]
			s++
			e--
			continue
		}

		if nums[s]%2 == 0 {
			s++
		}

		if nums[e]%2 == 1 {
			e--
		}

	}
	fmt.Println(nums)
	return nums
}
