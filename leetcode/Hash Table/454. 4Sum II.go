package Hash_Table
//四层循环会超时 可以先将两个值放入哈希表
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 创建一个哈希表
	m := make(map[int]int)

	for _,a := range nums1{
		for _,b := range nums2{
			//只取值
			m[a+b]++
		}
	}
	//a=1 b=1 m[2]=1
	count := 0
	for _,c := range nums3{
		for _,d := range nums4{
			//c=-2 d = 0 m[2] =1
			//a+b +c+d=0 ------ a+b = -c -d
			count += m[0-c-d]
		}
	}
	return count
}
