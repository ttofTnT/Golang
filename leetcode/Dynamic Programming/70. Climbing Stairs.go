package Dynamic_Programming
//爬到第一层楼梯有一种方法，爬到二层楼梯有两种方法。
//
//那么第一层楼梯再跨两步就到第三层 ，第二层楼梯再跨一步就到第三层。
//
//所以到第三层楼梯的状态可以由第二层楼梯 和 到第一层楼梯状态推导出来，那么就可以想到动态规划了。
func climbStairs(n int) int {
	if n==1 {
		return 1
	}
	dp := make([]int,n+1)
	//dp index max = n  0-n
	dp[1] = 1
	dp[2] = 2
	for i:=3 ; i<=n ;i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
