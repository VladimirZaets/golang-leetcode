package disjointset

type QuickUnion struct {
	roots []int
}

func NewQuickUnion(qty int) *QuickUnion {
	roots := make([]int, 10)
	for i, _ := range roots {
		roots[i] = i
	}
	return &QuickUnion{
		roots: roots,
	}
}

func (qu *QuickUnion) Find(index int) int {
	if qu.roots[index] == index {
		return index
	}
	return qu.Find(qu.roots[index])
}

func (qu *QuickUnion) Union(f int, s int) {
	froot := qu.Find(f)
	sroot := qu.Find(s)
	if sroot != froot {
		qu.roots[sroot] = froot
	}
}
