package graph

func InvertTree(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	var current *TreeNode

	for len(q) > 0 {
		current = q[0]
		q = q[1:]

		l := current.Left
		current.Left = current.Right
		current.Right = l

		if current.Left != nil {
			q = append(q, current.Left)
		}

		if current.Right != nil {
			q = append(q, current.Right)
		}
	}

	return root
}
