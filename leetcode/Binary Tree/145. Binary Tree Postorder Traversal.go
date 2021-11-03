package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode){
		if node == nil{
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res,node.Val) //这就是中
	}
	traversal(root)
	return res
}

