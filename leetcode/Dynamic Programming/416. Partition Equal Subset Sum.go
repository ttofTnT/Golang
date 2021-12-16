package Dynamic_Programming
func canPartition(nums []int) bool {
	sum :=0
	for _, value := range nums {
		sum += value
	}

	if sum % 2 == 1{
		return false
	}

	target := sum / 2
	dp := make([]int, target + 1)

	//  先遍历物品 在遍历背包
	//  背包要逆序便利
	for _, value := range nums {
		for j := target; j >= value;j--{
			if dp[j] < dp[j-value] + value  {
				dp[j] = dp[j-value] + value
			}
		}
	}

	return dp[target] == target
}