package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) <1  {
		return nil
	}
	index := findMax(nums)
	nodeval := nums[findMax(nums)]

	root := &TreeNode{
		Val: nodeval,
		Left: constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
	return root
}

func findMax (nums []int) int {
	index:=0
	for i:=0 ;i<len(nums);i++ {
		if nums[i] > nums[index] {
			index =i
		}
	}
	return index
}
