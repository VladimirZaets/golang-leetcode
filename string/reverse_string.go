package main

func reverseString(s []byte) {
	j := len(s) - 1
	h := len(s) / 2

	for i := 0; i < h; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}
