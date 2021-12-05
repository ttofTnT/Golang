package second_week

  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode

 }


func isSymmetric(root *TreeNode) bool {
	if root ==nil {
		return true
	}
	return isSametree(Inverttree(root.Left),root.Right)
}
func Inverttree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	Inverttree(node.Left)
	Inverttree(node.Right)

	node.Left ,node.Right = node.Right , node.Left
	return node
}

func isSametree(p,q *TreeNode) bool {
	if  p==nil && q == nil{
		return true
	}else if p!=nil && q!= nil {
		if p.Val != q.Val {
			return false
		}
		return isSametree(p.Left,q.Left) && isSametree(p.Right,q.Right)
	}else {
		return false
	}
}
