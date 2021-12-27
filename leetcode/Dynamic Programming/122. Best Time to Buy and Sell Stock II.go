package Dynamic_Programming

func maxProfit2(prices []int) int {
	//创建二维数组 0代表持有 1代表未持有
	dp := make([][]int,len(prices))

	for i := range dp {
		dp[i] = make([]int,2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i:=1 ; i < len(prices);i++{
		dp[i][0] = max(dp[i-1][0],dp[i-1][1] - prices[i])
		dp[i][1] = max(dp[i-1][0] + prices[i],dp[i-1][1])
	}
	return dp[len(prices)-1][1]
}

