package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil{
		return nil
	}
	nullnode := &TreeNode{
		Val: 0,
		Left: nil,
		Right: nil,
	}
	if root1 == nil {
		root1 = nullnode
	}
	if root2 == nil {
		root2 = nullnode
	}
	nodeval := root1.Val + root2.Val

	root := &TreeNode{
		Val: nodeval,
		Left: mergeTrees(root1.Left,root2.Left),
		Right: mergeTrees(root1.Right,root2.Right),
	}
	return root
}
