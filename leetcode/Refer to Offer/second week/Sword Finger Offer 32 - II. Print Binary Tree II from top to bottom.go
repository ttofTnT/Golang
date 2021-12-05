package second_week

func levelOrder2(root *TreeNode) [][]int {
	if root ==nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue,root)
	curNum , nextNum := 1 ,0
	temp := []int{}
	res := [][]int{}

	for len(queue) !=0 {
		for curNum >0 {
			node := queue[0]
			if node.Left !=nil {
				nextNum ++
				queue = append(queue, node.Left)
			}
			if node.Right !=nil {
				nextNum ++
				queue = append(queue, node.Right)
			}

			curNum--
			temp = append(temp,node.Val)
			queue = queue[1:]
		}

		curNum = nextNum
		nextNum = 0
		res = append(res, temp)
		temp = []int{} //清空 temp
	}
	return res
}