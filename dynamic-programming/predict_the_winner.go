package dynamic_programming

import "fmt"

func PpredictTheWinner(nums []int) bool {
	p := getBestTurn(nums, "F")
	return p["F"] >= p["S"]

}

func getBestTurn(nums []int, t string) map[string]int {
	q := t
	if q == "F" {
		q = "S"
	} else {
		q = "F"
	}

	if len(nums) == 1 {
		return map[string]int{t: nums[0]}
	}

	p1 := getBestTurn(nums[1:], q)
	p1[t] += nums[0]
	p2 := getBestTurn(nums[:len(nums)-1], q)
	p2[t] += nums[len(nums)-1]

	if p1[t] < p2[t] {
		return p2
	}
	return p1
}

func PredictTheWinner(nums []int) bool {
	r := GetBestTurn(nums, 0, len(nums)-1, 0)
	return r >= 0
}

func GetBestTurn(nums []int, l int, r int, i int) int {
	if l == r {
		return nums[l]
	}
	i++
	left := nums[l] - GetBestTurn(nums, l+1, r, i)
	right := nums[r] - GetBestTurn(nums, l, r-1, i)
	fmt.Println(left, right, i)

	if left > right {
		return left
	}
	return right
}
