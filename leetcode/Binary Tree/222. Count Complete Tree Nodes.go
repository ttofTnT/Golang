package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
递归遍历 左右中的后序遍历
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Left != nil {
		res += countNodes(root.Left)
	}
	if root.Right != nil {
		res += countNodes(root.Right)
	}

	return res
}

*/

//层序遍历来做

func countNodes(root *TreeNode) int {
	if root==nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	curNum, nextLevelNum , res := 1 , 0 ,1
	for len(queue) != 0 {
		node := queue[0]
		if node.Left !=nil {
			queue = append(queue, node.Left)
			nextLevelNum++
			res++
		}
		if node.Right !=nil {
			queue = append(queue,node.Right)
			nextLevelNum++
			res++
		}
		curNum--
		queue = queue[1:]
		if curNum ==0 {
			curNum = nextLevelNum
			nextLevelNum = 0
		}
	}
	return res
}
