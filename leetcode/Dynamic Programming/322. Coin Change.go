package Dynamic_Programming

func coinChange(coins []int, amount int) int {
	dp := make([]int , amount+1)
	//初始化 因为最小硬币是1 所以 coins[amount] <= amount
	for j := 1; j <= amount; j++ {
		dp[j] = amount + 1
	}
	dp[0] = 0

	//遍历物品
	for i:= 0 ; i< len(coins) ; i++ {
		//遍历背包
		for j:= coins[i] ; j<= amount ; j++{
			if dp[j-coins[i]] != amount+1  {
				dp[j] = min322(dp[ j- coins[i]] + 1,dp[j])
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min322(a, b int) int {
	if a < b {
		return a
	}
	return b
}