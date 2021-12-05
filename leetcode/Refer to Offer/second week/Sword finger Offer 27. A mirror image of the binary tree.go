package second_week

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	mirrorTree(root.Left)
	mirrorTree(root.Right)
	//调用在前 这样子先判断
	root.Left , root.Right = root.Right ,root.Left


	return root
}

