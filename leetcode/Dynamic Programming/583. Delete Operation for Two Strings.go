package Dynamic_Programming
func minDistance(word1 string, word2 string) int {
	//最大删除其实就是 找出两个数最长子序列长度

	m,n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
	}

	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			if word1[i-1] == word2[j-1]{
				dp[i][j] = dp[i-1][j-1] +1
			}else{
				dp[i][j] = max(dp[i-1][j],dp[i][j-1])
			}
		}
	}
	return m+n-dp[m][n]*2
}

//第二种 删除来做dp

func minDistance2(word1 string, word2 string) int {
	//dp[i][j] i-1结尾的word1 和 以j-1结尾的word2 相同最小步数

	m,n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range dp[0]{
		dp[0][j] = j
	}
	for i:=1; i<=m; i++{
		for j:=1; j<=n; j++{
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else {
				//删掉word1 或 删掉 word2 或两者删掉
				dp[i][j] = min(min(dp[i-1][j]+1,dp[i][j-1]+1),dp[i-1][j-1]+2)
			}
		}
	}
	return dp[m][n]
}