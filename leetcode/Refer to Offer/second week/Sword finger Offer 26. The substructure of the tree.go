package second_week

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if dfs(A,B) {
		return true
	}
	return isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
}

func dfs(A,B *TreeNode) bool{
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val{
		return false
	}
	return dfs(A.Left , B.Left)  && dfs(A.Right , B.Right)
}



