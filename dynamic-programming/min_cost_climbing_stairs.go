package dynamic_programming

func MinCostClimbingStairs(cost []int) int {
	first := MMinCostClimbingStairs(cost, 0, map[int]int{})
	second := MMinCostClimbingStairs(cost, 1, map[int]int{})
	if first < second {
		return first
	}
	return second
}

func MMinCostClimbingStairs(cost []int, start int, memo map[int]int) int {
	if _, ok := memo[start]; ok {
		return memo[start]
	}

	if start+2 < len(cost) {
		first := cost[start] + MMinCostClimbingStairs(cost, start+1, memo)
		second := cost[start] + MMinCostClimbingStairs(cost, start+2, memo)
		if first < second {
			memo[start] = first
			return memo[start]
		}
		memo[start] = second
		return memo[start]
	}
	return cost[start]
}
