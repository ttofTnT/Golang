package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	// BFS 层序遍历

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0{
		nextnode := []*TreeNode{}
		for _,node := range queue{
			if node.Left != nil{
				nextnode = append(nextnode, node.Left)
			}
			if node.Right != nil{
				nextnode = append(nextnode, node.Right)
			}
		}
		if len(nextnode) == 0 {
			return queue[0].Val //最后一行第一个
		}
		queue = nextnode
	}
}


