package second_week

func fib(n int) int {
	const mod  int = 1e9+7
	if n < 2 {
		return n
	}
	a, b, c := 0, 1, 1
	for i := 2; i < n; i++ {
		a, b = b, c
		c = (a + b) %  mod
	}
	return c
}
