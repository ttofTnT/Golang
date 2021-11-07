package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if(root==nil) {
		return 0
	}
	//左右中
	leftsum := sumOfLeftLeaves(root.Left)
	rightsum := sumOfLeftLeaves(root.Right)
	midsum := 0
	//中
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil  {
		midsum = root.Left.Val
	}
	sum := leftsum +rightsum +midsum
	return sum
}



