package third_week

/*
C++ 贪心
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int low = INT_MAX;
        int result = 0;
        for (int i = 0; i < prices.size(); i++) {
            low = min(low, prices[i]);  // 取最左最小价格
            result = max(result, prices[i] - low); // 直接取最大区间利润
        }
        return result;
    }
};
*/


func maxProfit(prices []int) int {
	length:=len(prices)
	if length==0{
		return 0
	}

	dp:=make([][]int,length)

	for i:=0;i<length;i++{
		dp[i]=make([]int,2)
	}

	dp[0][0]=-prices[0]
	dp[0][1]=0

	for i:=1;i<length;i++{
		dp[i][0]=max(dp[i-1][0],-prices[i])
		dp[i][1]=max(dp[i-1][1],dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

func max(a,b int)int {
	if a>b{
		return a
	}
	return b
}