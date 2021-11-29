package Refer_to_Offer

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, exists := m[num]; exists {
			return num
		}
		m[num] = true
	}
	return -1
}