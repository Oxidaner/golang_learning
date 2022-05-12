package main

func add(x, y int) (z int) {
	defer func() { //通过闭包延迟执行 // defer 延迟调用通过闭包读取和修改。
		z /= 2
	}()

	z = x + y
	return
}

func main() {
	println(add(2, 2))
}
