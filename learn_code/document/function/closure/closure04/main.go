package main

import "fmt"

// 外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

//在汇编层 ，test 实际返回的是 FuncVal 对象，其中包含了匿名函数地址、闭包对象指针。
//当调 匿名函数时，只需以某个寄存器传递该对象即可。
func main() {
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))
}
