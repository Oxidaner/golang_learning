package main

import "fmt"

func main() {
	//var a *int
	//*a = 100
	//
	//fmt.Println(*a)
	//
	//var b map[string]int
	//b["测试"] = 100
	//fmt.Println(b)

	/*
		new 初始化指针变量
	*/
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false

	/*
		先声明再初始化
	*/
	var aa *int
	aa = new(int)
	*aa = 10
	fmt.Println(*aa)
	fmt.Println(aa)

	/*
		生命并且初始化
	*/
	var m = new(int)
	*m = 100
	fmt.Println(m)
	fmt.Println(*m)
}
