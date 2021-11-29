package Refer_to_Offer

type MinStack struct {
	stack , minstack [] int
	index int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return  MinStack{
		stack: make([]int,0),
		minstack: make([]int,0),
		index: 0,
	}
}


func (this *MinStack) Push(x int)  {
	this.stack =append(this.stack,x)
	if this.index == 0 {
		this.minstack = append(this.minstack,x)
	}else {
		min := this.minstack[this.index-1]
		if x < min {
			this.minstack =append(this.minstack,x)
		}else {
			this.minstack =append(this.minstack,min)
		}
	}
	this.index++
}


func (this *MinStack) Pop()  {
	this.index--
	this.minstack = this.minstack[:this.index]
	this.stack = this.stack[:this.index]
}


func (this *MinStack) Top() int {
	return this.stack[this.index-1]
}


func (this *MinStack) Min() int {
	return this.minstack[this.index-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */