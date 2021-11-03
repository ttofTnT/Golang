package Binary_Tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode){
		if node == nil{
			return
		}
		traversal(node.Left)
		res = append(res,node.Val) //这就是中
		traversal(node.Right)
	}
	traversal(root)
	return res
}

