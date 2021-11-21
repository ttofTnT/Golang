package Backtracking

func partition(s string) [][]string {
	var path []string
	var res [][]string
	backtracking131(&res,path,0,s)
	return res
}
//index 是切割的下标
func backtracking131 (res *[][]string, path []string, index int,s string) {
	if len(s) == index {
		temp := make([]string,len(path))
		copy(temp,path)
		*res = append(*res, temp)
	}
	for i:= index; i < len(s); i++{
		if judge131(s,index,i) {
			path = append(path, s[index:i+1])//这里要i+1 切片的 取值
		}else {
			continue
		}

		backtracking131(res,path,i+1,s)

		path = path[:len(path)-1]
	}
}
func judge131(s string,left ,right int) bool{
	for ;left< right;  {
		if s[left] != s[right]{
			return false
		}
		left++
		right--
	}
	return true
}