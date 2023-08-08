package dynamic_programming

import "fmt"

func MaxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	tmp := make([]int, len(nums))
	for i, v := range nums {
		if i-1 >= 0 {
			tmp[i] = tmp[i-1] + v
		} else {
			tmp[i] = v
		}
	}
	return maxProfit(tmp)
}

//[-2 -1 -4 0 -1 1 2 -3 1]

func maxProfit(nums []int) int {
	minPoint := 0
	p := nums[minPoint]
	current := 1
	fmt.Println(nums)
	for current < len(nums) {
		sum := nums[current] - nums[minPoint]
		if sum > p {
			p = sum
		}
		if nums[current] > p {
			p = nums[current]
		}
		if nums[minPoint] > nums[current] {
			minPoint = current
		}
		current++
	}
	return p
}
