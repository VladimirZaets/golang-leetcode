package dynamic_programming

func CheckValidString(s string) bool {
	min := 0
	max := 0

	for _, v := range s {
		l := string(v)
		if l == "(" {
			min++
			max++
		}
		if l == ")" {
			min--
			max--
		}
		if l == "*" {
			min--
			max++
		}

		if max < 0 {
			return false
		}

		if min < 0 {
			min = 0
		}
	}

	if min != 0 {
		return false
	}

	return true
}
