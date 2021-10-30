package Double_pointer
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//快慢指针
	//设置 2 个指针，⼀个指针距离前⼀个指针 n 个距离。同时移动 2 个指
	//针，2 个指针都移动相同的距离。当⼀个指针移动到了终点，那么前⼀个指针就是倒数第 n 个节点了。
	dummy := &ListNode{
		Next: head,
	}
	fast := dummy
	low := dummy
	i:= 0
	for fast != nil {
		fast = fast.Next
		if i > n{
			low = low.Next
		}
		i++
	}
	low.Next = low.Next.Next
	return dummy.Next
}
