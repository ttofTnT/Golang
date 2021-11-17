package Backtracking


func combinationSum3(k int, n int) [][]int {
	var track []int //遍历路径
	var result [][]int //存放结果
	backtracking216(n,k,1,&track,&result)
	return result
}


// 个数为 k  和 为 n
func backtracking216(n,k, startindex int ,track*[]int ,result *[][]int) {
	if  len(*track) == k{
		var sum int
		temp := make([]int , k)
		copy(temp,*track)
		for _,v := range  *track{
			sum += v
		}
		//符合条件
		if sum == n {
			*result = append(*result,temp)
		}

		return
	}
	//len(track)+n-start+1 先不减枝
	for i:=startindex;i<=9; i++ {
		*track = append(*track,i)
		backtracking216(n,k,i+1,track,result)
		*track = (*track)[:len(*track)-1]
	}
}