package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 中序遍历 可以成一个有序数组
func isValidBST(root *TreeNode) bool {
	res := []int{}
	inOrder(root , &res)
	for i:=1 ; i< len(res) ; i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}
func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}
