package Greedy_Algorithm

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