package disjointset

type QuickFind struct {
	roots []int
}

func NewQuickFind(qty int) *QuickFind {
	roots := make([]int, qty)
	for i, _ := range roots {
		roots[i] = i
	}
	return &QuickFind{
		roots: roots,
	}
}

func (qf *QuickFind) Find(index int) int {
	return qf.roots[index]
}

func (qf *QuickFind) Union(f int, s int) {
	froot := qf.Find(f)
	sroot := qf.Find(s)
	for i, _ := range qf.roots {
		if qf.Find(i) == sroot {
			qf.roots[i] = froot
		}
	}
}

func (qf *QuickFind) GetRoots() []int {
	return qf.roots
}
