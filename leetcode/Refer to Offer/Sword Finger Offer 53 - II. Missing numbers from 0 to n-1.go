package Refer_to_Offer
func missingNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}

	return n
}
