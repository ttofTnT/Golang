package Binary_Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//左右孩子都为空的节点才是叶子节点
func minDepth(root *TreeNode) int {
	return getDepth(root)
}
func min(x,y int) int{
	if x > y {
		return y
	} else {
		return x
	}
}
func getDepth(node *TreeNode)int{
	if node ==nil {
		return 0
	}

	leftdepth , rightdepth := getDepth(node.Left) , getDepth(node.Right) //左右

	//中 后序遍历
	//左数空 右不空
	if node.Left == nil && node.Right !=nil{
		return rightdepth +1
	}
	//同理
	if node.Right ==nil && node.Left !=nil{
		return leftdepth + 1
	}

	return min(leftdepth,rightdepth) +1
}
