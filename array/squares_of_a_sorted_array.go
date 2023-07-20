package array

import "fmt"

func SortedSquaresFirst(nums []int) []int {
	var negatives []int
	pointer := 0

	for _, v := range nums {
		squared := v * v

		if v < 1 {
			negatives = append([]int{squared}, negatives...)
		} else {
			for pointer < len(negatives) && squared > negatives[pointer] {
				pointer += 1
			}

			negatives = append(negatives[:pointer], append([]int{squared}, negatives[pointer:]...)...)
			pointer += 1
		}
	}

	return negatives
}

func SortedSquares3(nums []int) []int {
	result := make([]int, len(nums))
	startPointer := 0
	endPointer := len(nums) - 1
	filled := len(result) - 1

	for endPointer >= len(nums)/2 && startPointer <= len(nums)/2 {
		startSqueared := nums[startPointer] * nums[startPointer]
		endSqueared := nums[endPointer] * nums[endPointer]
		if endPointer == startPointer {
			result[filled] = startSqueared
		} else if startSqueared >= endSqueared {
			result[filled] = startSqueared
			filled -= 1
			result[filled] = endSqueared
			filled -= 1
		} else {
			result[filled] = endSqueared
			filled -= 1
			result[filled] = startSqueared
			filled -= 1
		}
		endPointer -= 1
		startPointer += 1
	}

	return result
}

func sortedSquares(nums []int) []int {
	i := len(nums) - 1
	result := make([]int, len(nums))
	left := 0
	right := i

	for i >= 0 {
		rightValue := nums[right] * nums[right]
		leftValue := nums[left] * nums[left]

		if leftValue > rightValue {
			result[i] = leftValue
			left += 1
			i -= 1
		} else {
			result[i] = rightValue
			right -= 1
			i -= 1
		}
	}
	return result
}

func SquaresOfASortedArray(nums []int) []int {
	s := 0
	e := len(nums) - 1
	r := make([]int, len(nums))

	for s <= e {
		if nums[s]*nums[s] < nums[e]*nums[e] {
			r[e] = nums[s] * nums[s]
			s++
		} else {
			r[s] = nums[e] * nums[e]
			e--
		}
	}
	fmt.Println(r)
	return r
}
