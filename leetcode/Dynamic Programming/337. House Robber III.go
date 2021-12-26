package Dynamic_Programming


type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func rob3(root *TreeNode) int {
     res := robTree(root)
     return max(res[0], res[1])
}

func robTree(cur *TreeNode)[]int {//这个2维数组 0记录不偷最大值 和1记录偷了之后最大值
     if cur == nil {
           return []int{0,0}
     }

     //后序
     left := robTree(cur.Left)
     right := robTree(cur.Right)

     //  0 代表不偷 1代表偷
     //偷
     robCur := left[0] + right[0] + cur.Val
      //不偷
     nrobCur := max(left[0],left[1]) + max(right[0] , right[1])

     return []int{nrobCur,robCur}
}
