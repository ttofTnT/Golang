package Binary_Tree
//二叉搜索树中 求差值 最值 都可以用中序遍历 转化成 有序数据

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	var res []int
	findMIn(root,&res)

	min:=1000000//一个比较大的值
	for i:=1;i<len(res);i++{
		tempValue:=res[i]-res[i-1]
		if tempValue<min{
			min=tempValue
		}
	}
	return min
}
//中序遍历
func findMIn(root *TreeNode,res *[]int){
	if root==nil{
		return
	}
	findMIn(root.Left,res)
	*res=append(*res,root.Val)
	findMIn(root.Right,res)
}