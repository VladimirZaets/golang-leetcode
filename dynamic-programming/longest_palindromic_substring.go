package dynamic_programming

func LongestPalindrome(s string) string {
	res := ""

	for r := 0; r < len(s)+1; r++ {

		for l := 0; l < r; l++ {
			if r-l > len(res) && isPalindrome(s, l, r-1) {
				res = s[l:r]
			}
		}
	}

	return res
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
