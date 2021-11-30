package Greedy_Algorithm
//利润分解为每天为单位的维度，而不是从0天到第3天整体去考虑！
//假如第0天买入，第3天卖出，那么利润为：prices[3] - prices[0]。
// 相当于(prices[3] - prices[2]) + (prices[2] - prices[1]) + (prices[1] - prices[0])。
// 那么可以只加正利润

func maxProfit(prices []int) int {
	maxSum := 0
	for i:=1 ; i<len(prices) ; i++ {
		if prices[i] > prices[i-1] {
			maxSum += prices[i] - prices[i-1]
		}
	}
	return maxSum
}
