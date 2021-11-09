package Binary_Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//如果找到一个节点，发现左子树出现结点p，右子树出现节点q，或者 左子树出现结点q，右子树出现节点p，那么该节点就是节点p和q的最近公共祖先。
	if root == q || root == p || root == nil {
		return root
	}
	left := lowestCommonAncestor(root.Left,p,q)
	right := lowestCommonAncestor(root.Right,p,q)
	//如果搜索整个树，直接用一个变量left、right接住返回值，
	//这个left、right后序还有逻辑处理的需要，也就是后序遍历中处理中间节点的逻辑（也是回溯）。
	if left != nil && right!= nil {
		return root
	}
	if left == nil && right!=nil {
		return right
	}else if left !=nil && right==nil{
		return left
	}else {
		return nil

	}
}
