package string

import (
	"strings"
)

func ReverseWords(s string) string {
	sl := strings.Fields(s)
	ln := len(sl)
	reverse(sl, 0, ln-1)
	return strings.Join(sl, " ")
}

func reverse(sl []string, i int, j int) {
	for i <= j {
		sl[i], sl[j] = sl[j], sl[i]
		i++
		j--
	}
}
