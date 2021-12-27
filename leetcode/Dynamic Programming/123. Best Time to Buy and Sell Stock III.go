package Dynamic_Programming
func maxProfit3(prices []int) int {
	dp := make([][]int, len(prices))

	for i := range dp {
		dp[i] = make([]int, 5)
		// 五个状态
	}

	//初始化 0无操作 1 第一次买入 2 第一次卖出 3 第二次买入 4第二次卖出
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0

	for i:=1 ; i<len(prices) ; i++ {
		dp[i][0]=dp[i-1][0]
		dp[i][1]=max(dp[i-1][1],dp[i-1][0]-prices[i])	//	这次买入 和上次买入 对比
		dp[i][2]=max(dp[i-1][2],dp[i-1][1]+prices[i])	//  这次卖出（上次买入） 上次卖出对比
		dp[i][3]=max(dp[i-1][3],dp[i-1][2]-prices[i])
		dp[i][4]=max(dp[i-1][4],dp[i-1][3]+prices[i])
	}
	return dp[len(prices)][4]
}
