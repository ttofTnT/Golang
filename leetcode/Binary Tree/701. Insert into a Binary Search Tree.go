package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	//不破坏二叉搜索树 只需要在空结点插入就 行
	if root == nil{
		node := &TreeNode{
			val,
			nil,
			nil,
		}
		return node
	}
	//二叉搜索树性质
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
