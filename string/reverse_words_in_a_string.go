package main

import (
	"strings"
)

// func reverseWords(s string) string {
// 	si := -1
// 	var space uint8 = 32

// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] == space {
// 			if si == -1 {
// 				si = i
// 			} else {
// 				if s[i:si] == " " {
// 					s = s[:i] + s[si:]
// 				} else {
// 					s = s[:i] + s[si:] + s[i:si]
// 				}

// 				si = i
// 			}
// 		}
// 	}
// 	if si == -1 {
// 		return s
// 	}
// 	s = s[si+1:] + " " + s[:si]
// 	s = strings.Trim(s, " ")
// 	return s
// }

func reverseWords(s string) string {
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
