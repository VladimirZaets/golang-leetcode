package array

import "fmt"

func MoveZeroes(nums []int) {
	p := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			p++
			nums[p], nums[i] = nums[i], nums[p]
		}
	}
	fmt.Println(nums)
}
