package Dynamic_Programming
func findMaxForm(strs []string, m int, n int) int {
	//初始化
	dp := make([][]int , m+1)
	for i,_ := range dp {
		dp[i] = make([]int, n+1 )
	}

	//遍历
	for i := range strs {
		onum , znum := 0 , 0

		for _,v := range strs[i] {
			if v == '0' {
				znum++
			}
		}
		onum = len(strs[i]) - znum

		// 从后往前 遍历背包容量
		for j:= m ; j >= znum;j-- {
			for k:=n ; k >= onum;k-- {
				// 推导公式
				dp[j][k] = max474(dp[j][k],dp[j-znum][k-onum]+1)
			}
		}
	}
	return dp[m][n]
}

func max474(a,b int) int {
	if a > b {
		return a
	}
	return b
}