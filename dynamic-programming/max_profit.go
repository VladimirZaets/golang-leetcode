package dynamic_programming

func MaxProfit(prices []int) int {
	p := 0
	p1 := 0
	p2 := 1

	for p2 < len(prices) {
		if prices[p1] > prices[p2] {
			p1 = p2
		} else if prices[p1] < prices[p2] {
			c := prices[p2] - prices[p1]
			if c > p {
				p = c
			}
		}
		p2++
	}

	return p
}
