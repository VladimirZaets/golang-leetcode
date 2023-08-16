package dynamic_programming

func DDeleteAndEarn(nums []int) int {
	return ddeleteAndEarn(nums)
}

func DeleteAndEarn(nums []int) int {
	maxNumber := 0
	imap := map[int]int{}
	for _, v := range nums {
		imap[v] += v
		if maxNumber < imap[v] {
			maxNumber = imap[v]
		}
	}
	return dddeleteAndEarn(maxNumber, imap, map[int]int{})
}

func dddeleteAndEarn(v int, imap map[int]int, cache map[int]int) int {
	if _, ok := cache[v]; ok {
		return cache[v]
	}
	if v == 0 {
		return 0
	}

	if v == 1 {
		return imap[v]
	}

	gain := imap[v]
	first := dddeleteAndEarn(v-1, imap, cache)
	second := dddeleteAndEarn(v-2, imap, cache) + gain
	if first > second {
		cache[v] = first
	} else {
		cache[v] = second
	}

	return cache[v]
}

func ddeleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := 0

	for i, v := range nums {
		sum := 0
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		sl := removeFromSlice(tmp, i)
		sl = removePlusMinusOne(sl, v)
		sum = v + ddeleteAndEarn(sl)
		if sum > max {
			max = sum
		}
	}
	return max
}

func removePlusMinusOne(nums []int, v int) []int {
	needToDelete := []int{v + 1, v - 1}
	for i := 0; i < len(nums); i++ {
		if nums[i] == needToDelete[0] || nums[i] == needToDelete[1] {
			nums = removeFromSlice(nums, i)
			i = i - 1
		}
	}

	return nums
}

func removeFromSlice(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
