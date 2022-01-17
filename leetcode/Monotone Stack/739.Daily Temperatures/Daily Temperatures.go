package _39_Daily_Temperatures
func DailyTemperatures(temperatures []int) []int {
	//通常是一维数组，要寻找任一个元素的右边或者左边第一个比自己大或者小的元素的位置，此时我们就要想到可以用单调栈了。
	res := make([]int,len(temperatures))
	var mStack []int //定义一个单调栈
	for i,v := range temperatures{
		for len(mStack) >0 && temperatures[mStack[len(mStack)-1]] < v  { //栈不空   单调栈存储的是下标 遍历元素 大于栈顶元素
			//出栈 pop
			top := mStack[len(mStack)-1]
			mStack = mStack[:len(mStack)-1]

			res[top] = i - top
		}
		mStack = append(mStack,i)
	}
	return res
}
