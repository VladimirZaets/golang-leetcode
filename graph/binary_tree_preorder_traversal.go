package graph

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{root}
	var current *TreeNode

	for len(stack) != 0 {
		current = stack[len(stack)-1]
		result = append(result, current.Val)
		stack = stack[:len(stack)-1]
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return result
}
