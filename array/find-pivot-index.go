package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("The arguments must be passed. Example: 1,2,3,4,5")
		return
	}
	stingsSlice := strings.Split(flag.Args()[0], ",")
	intSlice := make([]int, len(stingsSlice))
	for i, v := range stingsSlice {
		intSlice[i], _ = strconv.Atoi(v)
	}

	fmt.Println(findPivotIndex(intSlice))

}

func findPivotIndex(nums []int) int {
	sum := 0
	sumLeft := 0
	sumWithoutCurrent := 0
	result := -1
	for _, v := range nums {
		sum += v
	}
	for i := 0; i < len(nums); i++ {
		sumWithoutCurrent = sum - nums[i]
		if sumLeft*2 == sumWithoutCurrent {
			result = i
			break
		}
		sumLeft += nums[i]

	}
	return result
}
