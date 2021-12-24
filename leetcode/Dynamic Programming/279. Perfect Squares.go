package Dynamic_Programming
func numSquares(n int) int {
	dp := make([]int , n+1)
	dp[0]= 0

	for i:=1 ; i <= n ; i++{
		dp [i] = n+1
	}
	//背包
	for i := 1 ; i <=n ; i++{
		//j是物品
		for j:= 1 ; j*j <=n ; j++ {
			if i >= j*j {
				dp[i] = min(dp[i-j*j]+1,dp[i])
			}
		}
	}
	return dp[n]
}

