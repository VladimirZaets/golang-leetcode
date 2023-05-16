package array

func ValidMountainArray(arr []int) bool {
	upCriteria := false
	downCriteria := false
	for i := 1; i < len(arr); i++ {
		if !downCriteria && arr[i] > arr[i-1] {
			upCriteria = true
		} else if upCriteria && arr[i] < arr[i-1] {
			downCriteria = true
		} else {
			return false
		}
	}

	if !upCriteria || !downCriteria {
		return false
	}

	return true
}

func ValidMountainArray2(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	var i int
	for i < len(arr)-1 && arr[i] < arr[i+1] {
		i++
	}
	if i == 0 || i == len(arr)-1 {
		return false
	}
	for i < len(arr)-1 && arr[i] > arr[i+1] {
		i++
	}

	return i == len(arr)-1
}
