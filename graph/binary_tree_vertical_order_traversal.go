package graph

import (
	"sort"
)

type IndexTreeNode struct {
	TreeNode *TreeNode
	Index    int
}

func VerticalOrder(root *TreeNode) [][]int {
	q := []*IndexTreeNode{&IndexTreeNode{root, 0}}
	mp := map[int][]int{0: []int{root.Val}}
	var current *IndexTreeNode

	for len(q) > 0 {
		current = q[0]
		q = q[1:]

		if _, ok := mp[current.Index]; !ok {
			mp[current.Index] = []int{}
		}

		if current.TreeNode.Left != nil {
			leftIndex := current.Index - 1
			if _, ok := mp[leftIndex]; !ok {
				mp[leftIndex] = []int{}
			}

			mp[leftIndex] = append(mp[leftIndex], current.TreeNode.Left.Val)
			q = append(q, &IndexTreeNode{current.TreeNode.Left, leftIndex})
		}

		if current.TreeNode.Right != nil {
			rignhtIndex := current.Index + 1
			if _, ok := mp[rignhtIndex]; !ok {
				mp[rignhtIndex] = []int{}
			}
			mp[rignhtIndex] = append(mp[rignhtIndex], current.TreeNode.Right.Val)
			q = append(q, &IndexTreeNode{current.TreeNode.Right, rignhtIndex})
		}
	}

	keys := make([]int, len(mp))
	i := 0

	for k := range mp {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	result := make([][]int, len(keys))
	i = 0

	for _, v := range keys {
		result[i] = mp[v]
		i++
	}

	return result
}
