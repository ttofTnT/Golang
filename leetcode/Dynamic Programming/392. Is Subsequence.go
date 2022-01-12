package Dynamic_Programming
func isSubsequence(s string, t string) bool {
	m , n := len(s) , len(t)
	if n < m {
		return false
	}

	dp := make([][]int, m+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
	}
	//dp[i][j] 代表表示以下标i-1为结尾的字符串s，和以下标j-1为结尾的字符串t，相同子序列的长度为dp[i][j]。
	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++{
			if s[i-1] == t[j-1]{
				dp[i][j] = dp[i-1][j-1]+1
			}else{
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	if dp[m][n] == m{
		return true
	}else {
		return false
	}
}

