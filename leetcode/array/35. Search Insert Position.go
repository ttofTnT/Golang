package array

func searchInsert(nums []int, target int) int {
	//二分法
	low , high := 0, len(nums)-1
	for low <= high {
		//右移一位，相当于除以2，但右移的运算速度更快
		//若使用(low+high)/2求中间位置容易溢出
		mid := low + (high-low)>>1
		if nums[mid] >target{
			high = mid - 1
		}else if nums[mid] < target{
			low = mid + 1
		}else {
			return mid
		}
		return high + 1
	}
}

