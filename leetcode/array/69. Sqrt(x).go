package array
//	求最后一个小于等于target的元素
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right, res := 1, x, 0
	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid < x/mid {
			left = mid + 1
			res = mid
			//此时的mid 是小于等于的
		} else if mid == x/mid {
			return mid
		} else {
			right = mid - 1
			//此时的mid 是大于等于的 所以可以一直循环到小于等于
		}
	}
	return res
}