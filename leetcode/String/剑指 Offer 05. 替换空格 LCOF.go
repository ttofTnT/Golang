package String
func replaceSpace(s string) string {
	b :=[]byte(s)
	//string不可变
	result := make([]byte,0)
	for i:=0; i<len(b) ;i++{
		if b[i] == ' ' {
			result = append(result,[]byte("%20")...) //解包 把byte[] 中的三个字符拆开 加到result里面
		}else {
			result = append(result,b[i])
		}
	}
	return string(result)
}

