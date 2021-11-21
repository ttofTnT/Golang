package Backtracking

func combine(n int, k int) [][]int {
	res := [][]int{}
	var backtracking77 func(start int, path []int)
	backtracking77 = func(start int, path []int) {
		if len(path) == k {
			//path存放单个答案
			temp := make([]int, len(path))
			copy(temp, path)
			//值拷贝 加到res后面 不改变
			res = append(res, temp)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			backtracking77(i+1, path)
			path = path[:len(path)-1]
			// 这里要返回 什么意思呢  4,2为例 当选了12时  要从2回去 到1 再选3
		}
	}
	backtracking77(1, []int{})
	return res
}
