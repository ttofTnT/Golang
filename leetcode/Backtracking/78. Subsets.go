package Backtracking

func subsets(nums []int) [][]int {
	res := [][]int{}
	var backtracking78 func(start int, path []int)
	backtracking78 = func(start int, path []int) {

		// 遍历所有树不用写条件
		temp := make([]int, len(path))
		copy(temp, path)
		//值拷贝 加到res后面 不改变
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking78(i+1, path)
			path = path[:len(path)-1]

		}
	}
	backtracking78(0, []int{})
	return res
}