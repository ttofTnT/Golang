package Binary_Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)
	//调用在前 这样子先判断
	 root.Left , root.Right = root.Right ,root.Left


	return root
}

