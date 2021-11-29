package Refer_to_Offer

 type Node struct {
    Val int
    Next *Node
    Random *Node
 }


func copyRandomList(head *Node) *Node {
	if head ==nil {
		 return head
	}

	mp := make(map[*Node]*Node) //map存储每个新节点存储
	for cur := head;cur != nil;cur = cur.Next {
		mp[cur] = NewNode(cur.Val)
	}

	for cur := head;cur != nil;cur = cur.Next {
		mp[cur].Next = mp[cur.Next]
		mp[cur].Random = mp[cur.Random]
	}

	return mp[head]
}

func NewNode (val int) *Node {
	return &Node{
		val,
		nil,
		nil,
	}
}
