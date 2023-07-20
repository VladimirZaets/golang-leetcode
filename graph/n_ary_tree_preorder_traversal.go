package graph

type Node struct {
	Val      int
	Children []*Node
}

func Preorder(root *Node) []int {
	result := []int{}
	stack := []*Node{root}
	var current *Node

	for len(stack) != 0 {
		current = stack[0]
		result = append(result, current.Val)
		stack = stack[1:]
		for i := len(current.Children) - 1; i >= 0; i-- {
			stack = append([]*Node{current.Children[i]}, stack...)
		}
	}

	return result
}
