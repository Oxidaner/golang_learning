package main

import "fmt"

type Mover interface {
	move()
}

type dog struct{}

//因为Go语言中有对指针类型变量求值的语法糖，
//dog指针fugui内部会自动求值*fugui。
func (d dog) move() {
	fmt.Println("狗会动")
}

//如果用指针接收者实现接口实现Mover接口的是*dog类型，
//所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值。
//func (d *dog) move() {
//	fmt.Println("狗会动")
//}

func main() {
	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	//var x Mover = dog{}
	x.move()
	var fugui = &dog{} // 富贵是*dog类型
	x = fugui          // x可以接收*dog类型
	x.move()
}

//一样的
