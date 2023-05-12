package disjointset

import (
	"fmt"
	"sort"
	"strings"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	uf := NewSmallestStringWIthSwaps(len(s))
	stringMap := make(map[int]string)
	ss := strings.Split(s, "")
	for i, v := range s {
		stringMap[i] = string(rune(v))
	}
	for _, v := range pairs {
		uf.Union(v[0], v[1])
	}
	branchesI := make(map[int][]int)
	for i, v := range uf.GetRoots() {
		if branchesI[v] == nil {
			branchesI[v] = []int{v}
		}

		if v != i {
			branchesI[v] = append(branchesI[v], i)
		}
	}
	for _, v := range branchesI {
		sort.Slice(v, func(i, j int) bool {
			return stringMap[v[i]] < stringMap[v[j]]
		})

		for i := 0; i < len(v); i += 2 {
			fmt.Println(v[i])
			tmp := ss[v[i]]
			ss[v[i]] = stringMap[v[i+1]]
			ss[v[i+1]] = tmp
		}
	}
	fmt.Println(branchesI)
	fmt.Println(stringMap)
	fmt.Println(s)
	fmt.Println(strings.Join(ss, ""))
	fmt.Println(uf.GetRoots())
	fmt.Println(ss)
	return strings.Join(ss, "")
}

type SmallestStringWIthSwaps struct {
	roots []int
}

func NewSmallestStringWIthSwaps(qty int) *SmallestStringWIthSwaps {
	roots := make([]int, qty)
	for i, _ := range roots {
		roots[i] = i
	}
	return &SmallestStringWIthSwaps{
		roots: roots,
	}
}

func (qf *SmallestStringWIthSwaps) Find(index int) int {
	return qf.roots[index]
}

func (qf *SmallestStringWIthSwaps) Union(f int, s int) {
	froot := qf.Find(f)
	sroot := qf.Find(s)
	for i, _ := range qf.roots {
		if qf.Find(i) == sroot {
			qf.roots[i] = froot
		}
	}
}

func (qf *SmallestStringWIthSwaps) GetRoots() []int {
	return qf.roots
}
