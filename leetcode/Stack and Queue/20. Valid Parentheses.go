package main



//golang  没有堆栈的函数
func isValid(s string) bool {

	if len(s)==0 {
		return true
	}

	stack := make([] rune,0)
	for _,v := range s{
		if v=='(' {
			stack = append(stack, ')')
		} else if v== '{' {
			stack =append(stack,'}')
		} else if v== '[' {
			stack = append(stack,']')
		} else if len(stack)>0 && stack[len(stack)-1] == v{
			stack = stack[:len(stack)-1]
		} else  {
			//左括号大于有括号 左括号不匹配右括号
			return false
		}
	}
	// 开始都是左括号没有有括号的情况
	return len(stack)==0
}
