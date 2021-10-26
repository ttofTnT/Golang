package Hash_Table

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n!= 1 && !m[n]{
		n,m[n] = getSum(n),true
		//用true 判断这个数是否得到过 如果得到过 就已经进入服了死循环
	}
	return  n == 1
}

func getSum(n int)int{
	sum := 0
	for n > 0 {
		sum += (n%10) * (n%10)
		n = n/10
	}
	return sum
}