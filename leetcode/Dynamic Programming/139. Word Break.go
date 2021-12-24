package Dynamic_Programming

func wordBreak(s string, wordDict []string) bool {
	//定义一个 map
	wordDictSet:=make(map[string]bool)
	for _,w:=range wordDict{
		//将单词 提取出来
		wordDictSet[w]=true
	}

	dp:=make([]bool,len(s)+1)
	dp[0]=true
	//背包
	for i:=1;i<=len(s);i++{
		//
		for j:=0;j<i;j++{
			if dp[j]&& wordDictSet[s[j:i]]{
				dp[i]=true
				break
			}
		}
	}
	return dp[len(s)]
}