package Backtracking

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	var histroy map [int]bool
	histroy = make(map[int]bool)
	sort.Ints(candidates)
	backtracking40(&res, path,0,0,candidates,target,histroy)
	return res
}

func backtracking40 (res *[][]int , path []int ,index,sum int, candidates []int, target int,  history map[int]bool) {
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
		if i>0 && candidates[i]==candidates[i-1]  && history[i-1]==false{
			continue
		}
		path = append(path,candidates[i])
		sum +=  candidates[i]
		history[i]=true
		backtracking40(res,path,i+1,sum,candidates,target,history)
		path =  path[:len(path)-1]
		sum -= candidates[i]
		history[i]=false
	}
}
