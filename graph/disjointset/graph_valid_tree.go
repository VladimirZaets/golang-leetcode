package disjointset

func ValidTree(n int, edges [][]int) bool {

	if n-len(edges) != 1 {
		return false
	}

	cities := NewQuickFind(n)

	for i := 0; i < len(edges); i++ {
		cities.Union(edges[i][0], edges[i][1])
	}

	keys := make(map[int]bool)
	unique := []int{}
	for _, v := range cities.GetRoots() {
		if _, value := keys[v]; !value {
			keys[v] = true
			unique = append(unique, v)
		}
	}

	return len(unique) == 1
}
