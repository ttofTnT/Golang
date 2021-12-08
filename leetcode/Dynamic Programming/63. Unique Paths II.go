package Dynamic_Programming

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 定义数据
	m , n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int,m)
	for i,_ := range dp{
		dp[i] = make([]int , n)
	}

	//初始化  边界条件
	for i:=0 ;i < m ; i++{
		if obstacleGrid[i][0] ==1 {
			break
		}
		dp[0][i] =1
	}

	for i:=0 ;i < n ; i++{
		if obstacleGrid[0][i] ==1 {
			break
		}
		dp[i][0] =1
	}
	//dp数组推到
	for  i :=1; i<m ; i++{
		for j := 1 ; j < n ; j++ {
			// 如果obstacleGrid[i][j]这个点是障碍物, 那么我们的dp[i][j]保持为0
			if obstacleGrid[i][j] != 1 {
				// 否则我们需要计算当前点可以到达的路径数
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
