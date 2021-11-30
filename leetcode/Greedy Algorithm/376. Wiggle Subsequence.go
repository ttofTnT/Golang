package Greedy_Algorithm

func wiggleMaxLength(nums []int) int {
	count := 1
	var curdif ,predif int
	if len(nums) < 2 {
		return count
	}

	for i:=0 ; i<len(nums)-1;i++ {
		curdif = nums[i+1] - nums[i]
		if predif * curdif < 0 || count == 1 && curdif != 0{ //第一次特殊情况
			predif = curdif
			count++
		}
	}
	return count
}