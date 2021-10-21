package array

//解法一 双指针 但会改变数组中的元素
func removeElement(nums []int, val int) int {
	length:=len(nums)
	res:=0
	for i:=0;i<length;i++{
		if nums[i]!=val {
			nums[res]=nums[i]
			res++
		}
	}
	return res
}

