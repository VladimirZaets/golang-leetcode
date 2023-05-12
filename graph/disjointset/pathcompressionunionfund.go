package disjointset

type PathCompressionUnionFind struct {
	roots []int
	ranks []int
}

func NewPathCompressionUnionFind(qty int) *PathCompressionUnionFind {
	roots := make([]int, qty)
	ranks := make([]int, qty)
	for i, _ := range roots {
		roots[i] = i
		ranks[i] = 0
	}
	return &PathCompressionUnionFind{
		roots: roots,
		ranks: ranks,
	}
}

func (pcuf *PathCompressionUnionFind) Find(index int) int {
	if pcuf.roots[index] == index {
		return index
	}
	pcuf.roots[index] = pcuf.Find(pcuf.roots[index])
	return pcuf.roots[index]
}

func (pcuf *PathCompressionUnionFind) Union(f int, s int) {
	froot := pcuf.Find(f)
	sroot := pcuf.Find(s)
	if sroot != froot {
		pcuf.roots[sroot] = froot
	}
	if pcuf.ranks[sroot] < pcuf.ranks[froot] {
		pcuf.roots[sroot] = froot
	} else if pcuf.ranks[sroot] > pcuf.ranks[froot] {
		pcuf.roots[froot] = sroot
	} else if pcuf.ranks[sroot] == pcuf.ranks[froot] {
		pcuf.roots[sroot] = froot
		pcuf.ranks[froot]++
	}
}

func (pcuf *PathCompressionUnionFind) GetRoots() []int {
	return pcuf.roots
}
