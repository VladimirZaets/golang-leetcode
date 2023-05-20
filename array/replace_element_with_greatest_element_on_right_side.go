package array

import "fmt"

func ReplaceElementWithGreatestElementOnRightSide(arr []int) []int {
	maxValue := arr[len(arr)-1]
	arr[len(arr)-1] = -1

	for i := len(arr) - 2; i >= 0; i-- {
		temp := arr[i]
		arr[i] = maxValue
		if temp > maxValue {
			maxValue = temp
		}
	}

	fmt.Println(arr)
	return arr
}
