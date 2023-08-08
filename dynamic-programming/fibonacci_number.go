package dynamic_programming

func Fib(n int) int {
	return FibMemo(n, map[int]int{})
}

func FibMemo(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}

	if n < 2 {
		return n
	}
	memo[n] = FibMemo(n-1, memo) + FibMemo(n-2, memo)
	return memo[n]
}
