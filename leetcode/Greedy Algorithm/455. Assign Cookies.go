package Greedy_Algorithm

import "sort"

func findContentChildren(g []int, s []int) int {
	//排序
	sort.Ints(g)
	sort.Ints(s)
	res:= 0
	index := len(s)-1 //饼干数组的下标
	// 从大到小
	for i := len(g)-1; i>=0  ;i-- {
		if index >=0  && s[index]>=g[i] {
			res++
			index--
		}
	}
	return res
}