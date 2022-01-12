package Dynamic_Programming

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	//不断调整起始位置
	for i := 1; i < len(nums);i++ {
		//如果之前num[i] 之前的值相加 小于0 就舍去
		if nums[i] + nums[i-1] >nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxSum{
			maxSum = nums[i]
		}
	}
	return maxSum
}

//dp

func maxSubArray2(nums []int) int {
	//dp[i]代表 以i结尾时候最大值和
	dp := make([]int,len(nums)+1)
	dp[1] = nums[0]
	for i:=2; i<= len(nums); i++ {
		dp[i] = max(nums[i-1],dp[i-1] + nums[i-1])
	}
	Max := dp[1]
	for i := 2; i <= len(nums); i++{
		if Max < dp[i]{
			Max = dp[i]
		}
	}
	return Max
}