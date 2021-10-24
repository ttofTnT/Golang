package List
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 双指针
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil{
		next := cur.Next
		// 保存了 下个节点 开始 翻转

		cur.Next = pre

		// 保存 已经翻转的节点
		pre = cur
		// 到 未翻转的节点
		cur = next
	}
	return pre
}
