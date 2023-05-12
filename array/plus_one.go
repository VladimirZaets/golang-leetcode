package array

func PlusOne(digits []int) []int {
	l := len(digits) - 1

	for {
		if l == -1 {
			digits = append([]int{1}, digits...)
			return digits
		}
		if digits[l] == 9 {
			digits[l] = 0
			l = l - 1
		} else {
			digits[l] = digits[l] + 1
			return digits
		}
	}
}
