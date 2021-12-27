package Dynamic_Programming
func maxProfit(prices []int) int {
	length:=len(prices)
	if length==0{
		return 0
	}

	dp:=make([][]int,length)

	for i:=0;i<length;i++{
		dp[i]=make([]int,2)
	}
	//[0]持有
	dp[0][0]=-prices[0]
	//[1]未持有
	dp[0][1]=0

	for i:=1;i<length;i++{
		//持有这次股票  这次买入 和 或上次最低价格买入
		dp[i][0]=max(dp[i-1][0],-prices[i])
		//未持有这次股票   这次卖掉 dp[i-1][0]+prices[i]  或者 这次也不卖  dp[i-1][1]
		dp[i][1]=max(dp[i-1][1],dp[i-1][0]+prices[i])
	}
	//最后未持有股票  一定是 价格最高的
	return dp[length-1][1]
}
