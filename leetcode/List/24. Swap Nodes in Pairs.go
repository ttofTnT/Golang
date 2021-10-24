package List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	// 虚拟头节点 注意&
	pre := dummy
	//head = list[i]   i：0-end
	//pre = list[i-1]
	for head != nil && head.Next != nil{
		//因为是交换 交换head.next 和 head  交换 list[i]  和 list[i+1]
		//此时 只要 list[i+1] 不空就行
		// 先保存节点
		next := head.Next.Next   // next 是下次开始的第一个节点

		// head 是  head.next 是 要交换的节点
		pre.Next = head.Next
		// pre 指向 交换后的 第一个节点
		head.Next.Next = head
		head.Next = next

		pre = head
		head = next

	}
	return dummy.Next
}