package main

import "fmt"

//剑指 Offer 48. 最长不含重复字符的子字符串
//请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//   请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


func lengthOfLongestSubstring(s string) int {
	//子字符串 是从i 开始 到start 结束
	lastOccurred := make(map[byte]int)
	fmt.Println(lastOccurred)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		fmt.Println(lastI)
		if ok && lastI>= start {
			start =lastOccurred[ch] +1
		}
		if i-start+1 > maxLength{
			maxLength = i -start+1
		}
		lastOccurred[ch] = i
	}
	fmt.Println(lastOccurred)
	return maxLength
}


func main()  {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring("bbbbb"))
	//fmt.Println(lengthOfLongestSubstring("pwwkew"))
}