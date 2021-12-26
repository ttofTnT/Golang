package Dynamic_Programming

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	// 成环 两种情况 1.包含 首地址 不含尾 和 2.包含尾地址 不含首
	result1 := robRange(nums, 0)
	result2 := robRange(nums, 1)
	return max(result1, result2)
}

// 偷盗指定的范围
func robRange(nums []int, start int) int {
	dp := make([]int, len(nums))
	dp[1] = nums[start]

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i - 2] + nums[i - 1 + start], dp[i - 1])
	}

	return dp[len(nums) - 1]
}