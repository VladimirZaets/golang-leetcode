package string

import (
	"strings"
)

func ReverseWordsInSting2(s string) string {
	slf := strings.Fields(s)
	for i, v := range slf {
		slf[i] = reverseString(v)
	}
	return strings.Join(slf, " ")
}

func reverseString(s string) string {
	b := []byte(s)
	j := len(s) - 1
	h := len(s) / 2

	for i := 0; i < h; i++ {
		b[i], b[j] = b[j], b[i]
		j--
	}

	return string(b)
}
