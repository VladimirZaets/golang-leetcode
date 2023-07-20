package array

func FindAllNumbersDisappearedInArray(nums []int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[nums[i]-1] {
			continue
		} else {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			i--
		}
	}

	for i, v := range nums {
		if i+1 != v {
			result = append(result, i+1)
		}
	}
	return result
}

func FindDisappearedNumbers(nums []int) []int {
	for i := range nums {
		index := abs(nums[i]) - 1

		if nums[index] > 0 {
			nums[index] *= -1
		}
	}
	res := []int{}
	for i := range nums {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
