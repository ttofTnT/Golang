package Refer_to_Offer

func search(nums []int, target int) int {
	left := helper(nums, 0, len(nums)-1, target)
	right := helper(nums, left, len(nums)-1, target+1)
	return right - left
}

func helper(nums []int, left, right, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}


