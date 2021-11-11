package Binary_Tree

//递归（隐含回溯）
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums)==0{return nil}//终止条件，最后数组为空则可以返回
	root:=&TreeNode{nums[len(nums)/2],nil,nil}//按照BSL的特点，从中间构造节点
	root.Left=sortedArrayToBST(nums[:len(nums)/2])//数组的左边为左子树
	root.Right=sortedArrayToBST(nums[len(nums)/2+1:])//数字的右边为右子树
	return root
}
