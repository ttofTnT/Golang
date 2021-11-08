package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <1 || len(inorder) <1 {
		return nil
	}
	nodeval := preorder[0]
	index := findNode(inorder,nodeval)

	root := &TreeNode{
		Val: nodeval,
		Left: buildTree( preorder[1:index+1] , inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1 :]),
	}

	return root

}

func findNode(inorder []int ,target int ) int {
	for i:=0;i<len(inorder);i++ {
		if target ==inorder[i] {
			 return i
		}
	}
	return -1
}

