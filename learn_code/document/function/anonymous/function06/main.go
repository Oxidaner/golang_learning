package main

func add(x, y int) (z int) {
	{ // 不能在一个级别，引发 "z redeclared in this block" 错误。
		z = x + y
		// return   // Error: z is shadowed during return
		return // 必须显式返回。
	}
}
func main() {
	print(add(1, 2))
}
