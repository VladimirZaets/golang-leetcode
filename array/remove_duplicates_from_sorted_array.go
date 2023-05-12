package array

func RemoveDuplicates(nums []int) int {
	prev := -101
	li := 0
	for i := 0; i < len(nums); i++ {
		if prev == nums[i] {
			nums[i] = -101
		} else {
			prev = nums[i]
			nums[li], nums[i] = nums[i], nums[li]
			li++
		}
	}
	return li
}
