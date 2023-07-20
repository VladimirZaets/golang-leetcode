package graph

func Postorder(root *Node) []int {
	stack := []*Node{root}
	result := []int{}
	var current *Node
	for len(stack) > 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append([]int{current.Val}, result...)
		for i := 0; i < len(current.Children); i++ {
			stack = append(stack, current.Children[i])
		}
	}
	return result
}
