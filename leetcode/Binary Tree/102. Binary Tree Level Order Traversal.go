package Binary_Tree

// 解法⼀ BFS
func levelOrder(root *TreeNode) [][]int {
	if root ==nil {
		return [][]int{}
	}
	//需要一个队列 关于*TreeNode的队列
	queue := []*TreeNode{}
	queue = append(queue,root)
	curNum , nextNum := 1 ,0
	temp := []int{} //保存每层节点
	res := [][]int{} //保存最后答案
 	//现在层数 节点个数 和 下一层节点个数
	for len(queue) !=0 {
		for curNum >0 {
			node := queue[0] // node 永远是第一个节点  经过切片
			if node.Left !=nil {
				nextNum ++
				queue = append(queue, node.Left)
			}
			if node.Right !=nil {
				nextNum ++
				queue = append(queue, node.Right)
			}
			// 已加载子节点 弹出头
			curNum--
			temp = append(temp,node.Val)
			queue = queue[1:]
		}
		//头节点全部弹出 重新规划头结点
		curNum = nextNum
		nextNum = 0
		res = append(res, temp)
		temp = []int{} //清空 temp
	}
	return res
}