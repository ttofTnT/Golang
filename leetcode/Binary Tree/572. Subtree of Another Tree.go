package Binary_Tree

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isSameTree(s, t) {
		return true
	}
	if s == nil {
		return false
	}
	if isSubtree(s.Left, t) || isSubtree(s.Right, t) {
		return true
	}
	return false
}

//func isSameTree(p *TreeNode, q *TreeNode) bool {
//	if p == nil && q == nil {
//		return true
//	} else if p != nil && q != nil {
//		if p.Val != q.Val {
//			return false
//		}
//		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
//	} else {
//		return false
//	}
//}