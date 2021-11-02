package main


func removeDuplicates(s string) string {
	//用栈实现 当先来的字符与栈顶一样时候两个都弹出
	stack := make([]rune,0)

	for _,v := range s{
		if len(stack)>0 && stack[len(stack)-1] == v {
			stack=stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}