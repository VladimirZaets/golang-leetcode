package graph

func InOrderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{root}
	var current *TreeNode

	for len(stack) > 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Left == nil && current.Right == nil {
			result = append(result, current.Val)
			continue
		}

		if current.Right != nil {
			stack = append(stack, current.Right)
			current.Right = nil
		}

		stack = append(stack, current)

		if current.Left != nil {
			stack = append(stack, current.Left)
			current.Left = nil
		}

	}
	return result
}
