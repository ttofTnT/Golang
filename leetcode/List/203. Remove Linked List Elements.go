package List
type ListNode struct {
	    Val int
	    Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	//虚拟头指针
	if head == nil{
		return head
	}
	//虚拟头节点
	newHead  := &ListNode{Val:0,Next: head}

	pre , cur := newHead , head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		}else  {
			pre= cur
		}
		cur = cur.Next
	}
	return newHead.Next
}