package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	p3 := &person{}            //使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	//p3.name = "博客"             //这是Go语言帮我们实现的语法糖 一下两行代码相同
	(*p3).name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}
}
