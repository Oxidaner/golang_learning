package main

import "fmt"

func main() {
	/*
	  - 程序定义一个int变量num的地址并打印
	  - 将num的地址赋给指针ptr，并通过ptr去修改num的值
	*/
	num := new(int)
	fmt.Println(num)

	ptr := num
	*ptr = 10
	println(*ptr)

	/*
		官方答案是
	*/
	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)
}
