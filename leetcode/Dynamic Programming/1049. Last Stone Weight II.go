package Dynamic_Programming
func lastStoneWeightII(stones []int) int {
	//30 * 100 /2 =1500+1
	dp := make([]int,1501)
	sum := 0

	for _,v := range stones {
		sum += v
	}

	target := sum / 2

	for _, v := range stones{
		for j:= target; j>=v; j-- {
			dp[j] = max1049(dp[j], dp[j-v] + v)
		}
	}
	return sum - dp[target]*2
}

func max1049(a, b int) int {
	if a > b {
		return a
	}
	return b
}
