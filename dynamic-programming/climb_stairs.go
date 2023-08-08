package dynamic_programming

func ClimbStairs(n int) int {
	return ClimbStairsMemo(n, map[int]int{})
}

func ClimbStairsMemo(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}

	if n == 2 {
		return 2
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	memo[n] = ClimbStairsMemo(n-1, memo) + ClimbStairsMemo(n-2, memo)

	return memo[n]
}
