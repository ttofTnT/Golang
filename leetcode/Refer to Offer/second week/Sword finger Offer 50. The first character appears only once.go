package second_week

func firstUniqChar(s string) byte {
	mp := [26]int{}
	for _, v := range s {
		mp[v-'a']++
	}
	for i, v := range s {
		if mp[v-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

