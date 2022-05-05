package main

import "fmt"

//类型定义: NewInt和 int不是同一个类型,但是有 int类型的特性
type NewInt int

//类型别名: MyInt和 int是一个类型
type MyInt = int

func main() {
	var a NewInt
	var b MyInt
	a = 10
	fmt.Println("a =", a)
	fmt.Printf("a = %d \n", a)
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
}
