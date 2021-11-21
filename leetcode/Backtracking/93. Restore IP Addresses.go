package Backtracking

import "strconv"

func restoreIpAddresses(s string) []string {
	var res []string
	var path [] string
	backtracking93(&res,path,0,s )
	return res
}
func backtracking93(res *[]string, path []string, startindex int,s string) {
	if startindex == len(s) && len(path) == 4 {
		//这里path 分为 四个划分部分
		tmpIpString:=path[0]+"."+path[1]+"."+path[2]+"."+path[3]
		*res=append(*res,tmpIpString)
	}

	for i:=startindex; i<len(s);i++ {
		path:= append(path,s[startindex:i+1])
		//放入 截取的 一段 ip 放入一段字符串
		if i-startindex +1 <=3 && len(path) <= 4 && judge93(s,startindex,i) {
			//第一个判断 截取的数字最多三位数 path中截取的 个数 小于四个
			backtracking93(res,path,i+1,s)
		}else {
			return
		}

		path = path[:len(path)-1]
	}
}

func judge93(s string,startindex,end int) bool {
	//段位以0为开头的数字不合法
	//段位里有非正整数字符不合法
	//段位如果大于255了不合法
	value ,_ := strconv.Atoi(s[startindex:end+1])
	if end - startindex >0 && s[startindex]=='0' {
		return false
	}
	if value >255 {
		return false
	}
	return true
}