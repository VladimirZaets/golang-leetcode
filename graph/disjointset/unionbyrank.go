package disjointset

type UnionByRank struct {
	roots []int
	ranks []int
}

func NewUnionByRank(qty int) *UnionByRank {
	roots := make([]int, 10)
	ranks := make([]int, 10)
	for i, _ := range roots {
		roots[i] = i
		ranks[i] = 0
	}
	return &UnionByRank{
		roots: roots,
		ranks: ranks,
	}
}

func (ubr *UnionByRank) Find(index int) int {
	if ubr.roots[index] == index {
		return index
	}
	return ubr.Find(ubr.roots[index])
}

func (ubr *UnionByRank) Union(f int, s int) {
	froot := ubr.Find(f)
	sroot := ubr.Find(s)

	if ubr.ranks[sroot] < ubr.ranks[froot] {
		ubr.roots[sroot] = froot
	} else if ubr.ranks[sroot] > ubr.ranks[froot] {
		ubr.roots[froot] = sroot
	} else if ubr.ranks[sroot] == ubr.ranks[froot] {
		ubr.roots[sroot] = froot
		ubr.ranks[froot]++
	}
}
