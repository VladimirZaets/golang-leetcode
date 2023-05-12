// Source: https://leetcode.com/problems/duplicate-zeros/
package array

import "fmt"

func DuplicateZerosSimple(arr []int) {
	arrLength := len(arr)
	for arrLength > 0 {
		if arr[arrLength-1] == 0 {
			part := arr[arrLength:]
			partLength := len(part)
			replace := 0
			for i := 0; partLength > i; i++ {
				tmp := part[i]
				part[i] = replace
				replace = tmp
			}
		}
		arrLength--
	}
	fmt.Println(arr)
}

func DuplicateZerosComplex(arr []int) {
	arrLength := len(arr)
	additional := 0
	for _, num := range arr {
		if num == 0 {
			additional++
		}
	}
	total := arrLength + additional

	for i := arrLength - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if total-1 < arrLength {
				arr[total-1] = 0
			}
			total--
		}
		if total-1 < arrLength {
			arr[total-1] = arr[i]
		}
		total--
	}
	fmt.Println(arr)
}
