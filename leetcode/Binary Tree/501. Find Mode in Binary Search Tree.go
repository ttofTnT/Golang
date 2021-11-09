package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	count := 1 //计数器
	var  res []int
	var  travel func (node *TreeNode)
	var pre *TreeNode
	maxcount := 0
	travel = func(node *TreeNode) {
		if node==nil {
			return
		}
		travel(node.Left)

		if pre!= nil && pre.Val == node.Val {
			count++
		}else {
			count = 1
		}

		if count >= maxcount {
			if len(res) >0 && count > maxcount {
				res = []int{node.Val}
			}else {
				//count >= maxcount （len(res)==0 || count <maxcount） ==  len(res)==0  && count == maxcount
				res = append(res,node.Val)
			}
			maxcount =count
		}
		pre = node
		travel(node.Right)
	}
	travel(root)
	return res
}

