package array
//1. 查找第⼀个值等于给定值的元素
//2. 查找最后⼀个值等于给定值的元素
//3. 查找第⼀个⼤于等于给定值的元素
//4. 查找最后⼀个⼩于等于给定值的元素
func searchRange(nums []int, target int) []int {
		return  []int {searchFirstEqualElement(nums,target),
			searchLastEqualElement(nums,target)}
}


// ⼆分查找第⼀个与 target 相等的元素，时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			//此时 已经找到与target 相等的元素 现在只需要判断边界条件向前减
			if mid == 0 || nums[mid-1] != target{
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// ⼆分查找最后⼀个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // 找到最后⼀个与 target 相等的元素
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}
// ⼆分查找第⼀个⼤于等于 target 的元素，时间复杂度 O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] >= target {
			if (mid == 0) || (nums[mid-1] < target) { // 找到第⼀个⼤于等于 target 的元 素
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
// ⼆分查找最后⼀个⼩于等于 target 的元素，时间复杂度 O(logn)
func searchLastLessElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] <= target {
			if (mid == len(nums)-1) || (nums[mid+1] > target) { // 找到最后⼀个⼩于等于 target 的元素
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
