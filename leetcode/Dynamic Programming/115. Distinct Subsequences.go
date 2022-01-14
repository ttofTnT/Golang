package Dynamic_Programming

func numDistinct(s string, t string) int {
	//dp[i][j]：以i-1为结尾的s子序列中出现以j-1为结尾的t的个数为dp[i][j]
	m,n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
		dp[i][0] = i //删掉s
	}

	for i:=1; i<=m; i++{
		for j:=1; j<=n; j++{
			if s[i-1] == t[i-1]{
				dp[i][j] = dp[i-1][j-1] +dp[i-1][j]
			}else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}