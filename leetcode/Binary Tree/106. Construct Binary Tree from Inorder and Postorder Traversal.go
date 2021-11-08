package Binary_Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 主要是 后序 postorder  可以 知道根节点 为 最后一个
//  根据根节点 可以 通过 中序遍历 知道 前半部分
func buildTree6(inorder []int, postorder []int) *TreeNode {
	if len(inorder) <1 || len(postorder)<1 {
		return nil
	}

	nodeval := postorder[len(postorder)-1]
	index := findNode1(inorder,nodeval)

	root := &TreeNode{
		Val: nodeval,
		Left: buildTree6(inorder[:index], postorder[:index]),
		Right: buildTree6(inorder[index+1:],postorder[index:len(postorder)-1]),
		//左右节点不同遍历 不同 看仔细跳过 头节点
	}
	return root

}

func findNode1(t []int , target int) int {
	for i:=0;i<len(t);i++ {
		if t[i] == target {
			return i
		}
	}
	return -1
}