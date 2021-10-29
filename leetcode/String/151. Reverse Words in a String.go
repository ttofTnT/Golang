package String

func reverseWords(s string) string {
	left , right:= 0,0
	b := []byte(s) //将string转换成 byte[]
	//消除冗余 字符前面
	for len(b)>0 && right<len(b) && b[right] == ' ' {
		right++
	}
	//跳过空余的地区 后面 用right 给left 赋值

	//	消除中间冗余
	for ;right<len(b);right++ {
		if right-1>0 &&b[right] == b[right-1] && b[right] == ' ' {
			//right -1 大于0才能标准不越界
			continue
		}else {
			b[left] = b[right]
			left++
		}
	}
	//消除末尾冗余
	//这时候要看 left 这时候right 已经遍历完
	// 但会出现两种情况 原string 含有" " 和原 string 不含有空格
	if left-1>0 && b[left-1] == ' ' {
		b = b[:left-1]  //切片 [0,left-2]
	}else {
		b = b[:left]  //切片 [0,left-]
	}
	//进行翻转总翻转
	reverse(&b,0,len(b)-1)

	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)

}

func reverse(b *[]byte, left ,right int){
	for left < right{
		(*b)[left],(*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}


//import "strings"
//func reverseWords151(s string) string {
//	ss := strings.Fields(s)
//	reverse151(&ss, 0, len(ss)-1)
//	return strings.Join(ss, " ")
//}
//func reverse151(m *[]string, i int, j int) {
//	for i <= j {
//		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
//		i++
//		j--
//	}
//}