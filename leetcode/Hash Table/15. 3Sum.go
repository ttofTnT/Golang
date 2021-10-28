package Hash_Table



import "sort"

//说是哈希表主题 但是这道题 用双指针比较好
//i left = i+1 right = len()-1

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	// 先排序

	for i:=0;i<len(nums)-2;i++ {
		//	这个i遍历的值可以减2 因为不能重复 i 是最小值
		n1 := nums[i]
		if n1>0 {
			//最小的一个值大于0
			break
		}
		if i>0&&n1 == nums[i-1] {
			//去重复
			continue
		}
		r,l := len(nums)-1, i+1
		for r>l {
			//n2 ， n3 会改变的 必须要写在循环内
			n2 := nums[l]
			n3 := nums[r]
			res := n1 + n2 + n3
			if res == 0 {
				result= append(result,[]int{n1,n2,n3} )
				//还需要收缩 第一处去重复只去掉了第一个字符的重复
				for r>l && n2 == nums[l]{
					l++
				}
				for r>l && n3 == nums[r]{
					r--
				}
			}else if res > 0 {
				//太大了
				r--
			}else {
				l++
			}
		}
	}
	return result
}

