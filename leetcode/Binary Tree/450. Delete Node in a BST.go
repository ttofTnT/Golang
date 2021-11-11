package Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	//1.没找到节点
	if root ==nil {
		return root
	}
	//找到节点了
	if root.Val == key {
		//2.左右节点都为空
		// 后面root 的left 和right 都有值了 删除节点 返空就行
		if root.Left==nil && root.Right ==nil  {
			return nil
		}else if root.Left == nil {
			//到这里一定满足 左右孩子不同时为空
			//3.左孩子为空
			return root.Right
		} else if root.Right ==nil {
			//4.右孩子为空
			return root.Left
		} else {
			//左右孩子都有值
			// 左孩子放到 右孩子的最左孩子上  左孩子 <p 右孩子都>p  右孩子的最左孩子值为最小大于p
			cur := root.Right
			for cur.Left !=nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			root =root.Right
			return root
		}
		//新的节点返回给上一层，上一层就要用 root->left 或者 root->right接住
	}else if root.Val>key {
		root.Left = deleteNode(root.Left,key)
	}else {
		root.Right = deleteNode(root.Right,key)
	}
	return root

}