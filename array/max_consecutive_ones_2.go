package array

import "fmt"

func MaxConsecutiveOnes(nums []int) int {
	fliped := -1
	max := 0
	current := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			current++
		}
		if nums[i] == 0 && fliped == -1 {
			current++
			fliped = i
		} else if nums[i] == 0 && fliped != -1 {
			if current > max {
				max = current
			}
			current = i - fliped
			fliped = i
		}
	}

	if current > max {
		max = current
	}
	fmt.Println(max)
	return max
}

func FindMaxConsecutiveOnes2(nums []int) int {
	k := 1
	left := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			k--
		}

		if k < 0 {
			if nums[left] == 0 {
				k++
			}

			left++
		}
	}

	fmt.Println(len(nums) - left)
	return len(nums) - left
}
