package second_week

func levelOrder(root *TreeNode) []int {
	ans := []int{}
	que := []*TreeNode{root}

	if root == nil {
		return ans
	}

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		ans = append(ans, cur.Val)

		if cur.Left != nil {
			que = append(que, cur.Left)
		}
		if cur.Right != nil {
			que = append(que, cur.Right)
		}
	}

	return ans
}