package Backtracking

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	var backtracking90 func(start int, path []int)
	backtracking90 = func(start int, path []int) {

		// 遍历所有树不用写条件
		temp := make([]int, len(path))
		copy(temp, path)
		//值拷贝 加到res后面 不改变
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			if i>start &&nums[i]==nums[i-1]{
				continue
			}
			path = append(path, nums[i])
			backtracking90(i+1, path)
			path = path[:len(path)-1]

		}
	}
	backtracking90(0, []int{})
	return res
}