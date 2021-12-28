package Dynamic_Programming

func findLength(A []int, B []int) int {
	//初始化一个二维数组 dp【i】【j】 代表i-1 为结尾的  A   和 代表为j-1 为结尾的 B
	m := len(A)
	n := len(B)
	dp := make([][]int,m+1)
	res := 0
	for i := range dp {
		dp[i] = make([]int,n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}