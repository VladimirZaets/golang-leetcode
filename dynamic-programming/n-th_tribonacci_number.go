package dynamic_programming

func Tribonacci(n int) int {
	if n == 0 {
		return n
	}
	return TribRec(n, map[int]int{})
}

func TribRec(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if n == 1 {
		return 1
	}
	if n <= 3 {
		return n - 1
	}

	memo[n] = TribRec(n-1, memo) + TribRec(n-2, memo) + TribRec(n-3, memo)
	return memo[n]
}
