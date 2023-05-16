package array

func CheckIfNandItsDoubleExist(arr []int) bool {
	m := map[int]bool{arr[0]: true}
	for i := 1; i < len(arr); i++ {
		if m[arr[i]*2] || (arr[i]%2 == 0 && m[arr[i]/2]) {
			return true
		}
		m[arr[i]] = true
	}
	return false
}
