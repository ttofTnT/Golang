package Binary_Tree



//type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode){
		if node == nil{
			return
		}
		res = append(res,node.Val) //这就是中
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}
