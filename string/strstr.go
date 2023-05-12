package string

func StrStr(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)
	if haystack == needle {
		return 0
	}
	for i := 0; i <= haystackLen-needleLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}
