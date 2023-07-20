package graph

func LevelOrder(root *Node) [][]int {
	result := [][]int{}
	q := []*Node{root}
	i := 0
	ll := 0
	nextLl := 0
	var current *Node
	level := []int{}
	for len(q) > 0 {

		current = q[0]
		q = q[1:]

		if i == ll {
			ll = nextLl + len(current.Children)
			nextLl = ll
			level = append(level, current.Val)
			cpLevel := make([]int, len(level))
			copy(cpLevel, level)
			result = append(result, cpLevel)
			level = []int{}
		} else if i < ll {
			level = append(level, current.Val)
			nextLl = nextLl + len(current.Children)
		}

		if current.Children != nil {
			q = append(q, current.Children...)
		}
		i++
	}
	return result
}

func LevelOrder2(root *Node) [][]int {
	res, queue, levelRes, levelCount := [][]int{}, []*Node{root}, []int{}, 1
	if root == nil {
		return res
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		queue = append(queue, node.Children...)
		levelRes = append(levelRes, node.Val)
		levelCount--

		if levelCount == 0 {
			res = append(res, levelRes)
			levelRes = []int{}
			levelCount = len(queue)
		}
	}
	return res
}

//1
//3 2 4
//5 6
