package Hash_Table
//顺序扫描数组，对每⼀个元素，在 map 中找能组合给定值的另⼀半数字，如果找到了，直接返回 2 个
//数字的下标即可。如果找不到，就把这个数字存⼊ map 中，等待扫到“另⼀半”数字的时候，再取出来返
//回结果。

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}
