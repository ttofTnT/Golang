package Greedy_Algorithm

func canJump(nums []int) bool {
	cover := 0
	if len(nums) == 1 {//一个元素
		return true
	}
	for i:=0; i<=cover;i++ {//小于的是cover 挑战范围
		cover = maxint(i+nums[i],cover)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}
func maxint(i,j int)  int{
	if  i>=j{
		return i
	}else {
		return j
	}
}