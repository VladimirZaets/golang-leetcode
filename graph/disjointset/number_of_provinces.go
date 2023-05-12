package disjointset

func FindCircleNum(isConnected [][]int) int {
	cities := NewQuickFind(len(isConnected))

	for i := 0; i < len(isConnected); i++ {
		for j := i + 1; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 {
				cities.Union(i, j)
			}
		}
	}

	keys := make(map[int]bool)
	unique := []int{}
	for _, v := range cities.GetRoots() {
		if _, value := keys[v]; !value {
			keys[v] = true
			unique = append(unique, v)
		}
	}
	return len(unique)
}
