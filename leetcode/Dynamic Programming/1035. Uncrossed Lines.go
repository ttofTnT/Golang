package Dynamic_Programming

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	//1143. 最长公共子序列一模一样
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int,m+1)
	for i := range dp{
		dp[i] = make([]int,n+1)
	}
	dp[0][0] = 0
	//dp[m][n] 表示 m个 nums1 和 n个nums2 能达到的最长公共子序列长度
	for i:= 1; i <= m; i++ {
		for j:=1; j <=n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1]+1
			} else {
				dp[i][j] = max(dp[i-1][j],dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}


