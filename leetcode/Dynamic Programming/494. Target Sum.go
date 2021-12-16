package Dynamic_Programming

func findTargetSumWays(nums []int, target int) int {
	//记 + 号的⼀组 P ，记 - 号的⼀组 N
	//sum(P) - sum(N) = target
	//sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
	//2 * sum(P) = target + sum(nums)
	//那么这道题就转换成了，能否在数组中找到这样⼀个集合，和等于 (target + sum(nums)) / 2
	var sum int
	for _,value:=range nums{
		sum+=value
	}
	//如果sum<target或者 sum+target不是偶数（因为a是int） 或者两者之和小于0了
	if sum<target||(sum+target)%2==1||(sum+target)<0{
		return 0
	}
	//开始dp初始化
	dp:=make([][]int,len(nums)+1)
	for i:=0;i<=len(nums);i++{
		tmp:=make([]int,(target+sum)/2+1)//背包容量
		dp[i]=tmp
	}
	dp[0][0]=1//当背包容量为0，且物品为0时，填满背包就1种方法
	for i:=0;i<len(nums)+1;i++{
		if i==0{
			continue
		}
		for j:=0;j<(target+sum)/2+1;j++{
			if nums[i-1]<=j{//如果背包装的下
				dp[i][j]=dp[i-1][j]+dp[i-1][j-nums[i-1]]
			}else{
				dp[i][j]=dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][(target+sum)/2]
}