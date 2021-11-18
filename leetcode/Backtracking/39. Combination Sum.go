package Backtracking

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var  path []int
	backtracking39(&res, path,0,0,candidates,target)
	return res
}

func backtracking39 (res *[][]int , path []int ,index,sum int, candidates []int, target int ) {
	if  sum >target{
		return
	}
	if sum == target {
		temp:= make([]int,len(path))
		copy(temp,path)
		*res = append(*res, temp)
		return
	}

	for i:=index; i<len(candidates) ; i++ {
		path = append(path,candidates[i])
		sum +=  candidates[i]

		backtracking39(res,path,i,sum,candidates,target)
		path =  path[:len(path)-1]
		sum -= candidates[i]
	}
}
