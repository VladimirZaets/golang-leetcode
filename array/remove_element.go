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
