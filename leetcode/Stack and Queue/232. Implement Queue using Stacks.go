package main

type MyQueue struct {
	stack_in [] int
	stack_out [] int
}


func Constructor() MyQueue {
	return MyQueue{
		stack_in: make([]int,0),
		stack_out : make([]int,0),
	}
}


func (this *MyQueue) Push(x int)  {
	//进栈要看栈出有没有为空
	for len(this.stack_out) != 0 {
		val := this.stack_out[len(this.stack_out)-1]
		this.stack_out = this.stack_out[:len(this.stack_out)-1]
		this.stack_in = append(this.stack_in, val)
	}
	this.stack_in = append(this.stack_in, x)

}


func (this *MyQueue) Pop() int {
	//出栈 要先把 stack_in的 东西全部到stack_out
	for len(this.stack_in) != 0{
		val := this.stack_in[len(this.stack_in)-1] //获取最后一个
		this.stack_in =this.stack_in[:len(this.stack_in)-1]
		this.stack_out = append(this.stack_out,val )
	}
	if len(this.stack_out) == 0 {
		return 0
	}
	//在弹出
	val := this.stack_out[len(this.stack_out)-1]
	this.stack_out = this.stack_out[:len(this.stack_out)-1]
	return val
}


func (this *MyQueue) Peek() int {
	for len(this.stack_in) != 0{
		val := this.stack_in[len(this.stack_in)-1] //获取最后一个
		this.stack_in =this.stack_in[:len(this.stack_in)-1]
		this.stack_out = append(this.stack_out,val )
	}
	if len(this.stack_out) == 0 {
		return 0
	}
	//在弹出
	val := this.stack_out[len(this.stack_out)-1]
	return val
}


func (this *MyQueue) Empty() bool {
	return len(this.stack_out) == 0 && len(this.stack_in)== 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */