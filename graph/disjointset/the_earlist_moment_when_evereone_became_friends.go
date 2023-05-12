package disjointset

import (
	"sort"
)

func EarliestAcq(logs [][]int, n int) int {
	cities := NewQuickFind(n)
	i := 0

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	for i < len(logs) {
		cities.Union(logs[i][1], logs[i][2])
		if len(getUnique(cities.GetRoots())) == 1 {
			return logs[i][0]
		}
		i++
	}

	return -1
}

func getUnique(list []int) []int {
	keys := make(map[int]bool)
	unique := []int{}
	for _, v := range list {
		if _, value := keys[v]; !value {
			keys[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}
