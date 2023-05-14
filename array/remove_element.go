package array

func RemoveElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		nums[i] = -1
		if v != val {
			nums[k] = v
			k++
		}
	}
	return k
}

func RemoveElement2(nums []int, val int) int {
	e := len(nums) - 1
	replaced := 0
	for l := len(nums) - 1; l >= 0; l-- {
		if nums[l] == val {
			nums[l] = nums[e]
			replaced++
			nums[e] = -1
			e--
		}
	}
	return len(nums) - replaced
}
