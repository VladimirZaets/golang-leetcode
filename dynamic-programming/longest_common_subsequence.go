package dynamic_programming

import "fmt"

func LongestCommonSubsequence(text1 string, text2 string) int {
	var shorter string
	var longer string

	if len(text1) < len(text2) {
		shorter = text1
		longer = text2
	} else {
		shorter = text2
		longer = text1
	}

	return longestCommonSubsequenceDp(shorter, longer, len(shorter)-1, len(longer)-1, map[string]int{})
}

func longestCommonSubsequenceDp(shorter string, longer string, i int, j int, memo map[string]int) int {
	ij := fmt.Sprintf("%d:%d", i, j)
	if _, ok := memo[ij]; ok {
		return memo[ij]
	}
	if i == -1 {
		return 0
	}

	if j == -1 {
		return 0
	}
	if shorter[i] == longer[j] {
		memo[ij] = longestCommonSubsequenceDp(shorter, longer, i-1, j-1, memo) + 1
		return memo[ij]
	} else {
		p1 := longestCommonSubsequenceDp(shorter, longer, i-1, j, memo)
		p2 := longestCommonSubsequenceDp(shorter, longer, i, j-1, memo)
		if p1 > p2 {
			memo[ij] = p1
		} else {
			memo[ij] = p2
		}
		return memo[ij]
	}
}
