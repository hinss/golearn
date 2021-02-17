package fib

// 声明大写把这个方法公有化暴露出去
func Fbn() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, b+a
		return a
	}
}
