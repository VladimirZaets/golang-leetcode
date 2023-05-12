package array

import "strconv"

func FindNumbers(nums []int) int {
	evenQty := 0
	for _, v := range nums {
		s := strconv.Itoa(v)
		if len(s)%2 == 0 {
			evenQty += 1
		}
	}
	return evenQty
}
