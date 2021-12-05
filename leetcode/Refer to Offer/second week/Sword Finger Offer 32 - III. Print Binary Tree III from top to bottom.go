package second_week

func levelOrder3(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	var Level int = 0
	for len(queue) != 0 {
		temp := []*TreeNode{}
		res = append(res, []int{})
		for _, v := range queue {
			res[Level] = append(res[Level], v.Val)
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		if Level % 2 != 0 {
			Reverse(res[Level])
		}
		Level++
		queue = temp
	}
	return res
}

func Reverse(arr []int) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		arr[i],arr[length-1-i] = arr[length-1-i],arr[i]
	}
}

