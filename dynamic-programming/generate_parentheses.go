package dynamic_programming

import (
	"fmt"
	"strconv"
)

func GenerateParenthesis(n int) []string {
	c := n
	r := []string{}
	generateParenthesis(n-1, "(", c, &r)
	return r
}

func GenerateParenthesisMemo(n int) []string {
	c := n
	r := []string{}
	s := generateParenthesisMemo(n-1, "(", c, map[string]string{})
	fmt.Println(s)
	return r
}

func generateParenthesis(n int, s string, c int, r *[]string) {
	if n == 0 && c == 0 {
		*r = append(*r, s)
	}

	if n > 0 && c > 0 {
		generateParenthesis(n-1, s+"(", c, r)
	}

	if c > 0 && n != c {
		generateParenthesis(n, s+")", c-1, r)
	}
}

func generateParenthesisMemo(n int, s string, c int, memo map[string]string) string {
	//if _, ok := memo[strconv.Itoa(n)+strconv.Itoa(c)]; ok {
	//	return memo[strconv.Itoa(n)+strconv.Itoa(c)]
	//}

	if n == 0 && c == 0 {
		return s
	}

	if n > 0 && c > 0 {
		r := generateParenthesisMemo(n-1, s+"(", c, memo)
		memo[strconv.Itoa(n-1)+strconv.Itoa(c)] = r
		fmt.Println(r, n-1, s)
	}

	if c > 0 && n != c {
		r := generateParenthesisMemo(n, s+")", c-1, memo)
		memo[strconv.Itoa(n)+strconv.Itoa(c-1)] = r
		fmt.Println(r, n, c-1)

	}

	return ""
}

//n = 5
//["((((()))))","()()()()()", "(()()()())","(())()()()","()()()(())", "()()(())()", "()(())()()", "(())(())()", "()(())(())" ]

//n = 4
//["(((())))","()()()()", "(()()())","(())()()","()()(())", "()(())()", "(())(())" ]

//n = 3
//["((()))", "()()()", "(()())","(())()","()(())"]

//n = 2
//["(())","()()"]

//n = 1
//["()"]
