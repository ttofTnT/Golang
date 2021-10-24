package List

type MyLinkedList struct {
	Val int
	next *MyLinkedList
}


func Constructor() MyLinkedList {
	return MyLinkedList{Val: -999,next: nil}
}


func (this *MyLinkedList) Get(index int) int {
	cur := this
	if cur.Val == index {
		if cur.Val == -999 {
			return -1
		}else {

		}

	}

}


func (this *MyLinkedList) AddAtHead(val int)  {

}


func (this *MyLinkedList) AddAtTail(val int)  {

}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {

}


func (this *MyLinkedList) DeleteAtIndex(index int)  {

}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */