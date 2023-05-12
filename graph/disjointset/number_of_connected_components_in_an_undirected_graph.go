package disjointset

func countComponents(n int, edges [][]int) int {
	cities := NewNumberOfConnectedComponents(n)

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

	return len(unique)
}

type NumberOfConnectedComponents struct {
	roots []int
}

func NewNumberOfConnectedComponents(qty int) *NumberOfConnectedComponents {
	roots := make([]int, qty)
	for i, _ := range roots {
		roots[i] = i
	}
	return &NumberOfConnectedComponents{
		roots: roots,
	}
}

func (qf *NumberOfConnectedComponents) Find(index int) int {
	return qf.roots[index]
}

func (qf *NumberOfConnectedComponents) Union(f int, s int) {
	froot := qf.Find(f)
	sroot := qf.Find(s)
	for i, _ := range qf.roots {
		if qf.Find(i) == sroot {
			qf.roots[i] = froot
		}
	}
}

func (qf *NumberOfConnectedComponents) GetRoots() []int {
	return qf.roots
}
