package main

import "fmt"

//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 设置p的年龄
// 使用值接收者 会拷贝
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func main() {
	p1 := NewPerson("测试", 25)
	p1.Dream()

	fmt.Println(p1.age) // 25
	p1.SetAge(30)
	fmt.Println(p1.age) // 30
	p1.SetAge2(40)      // (*p1).SetAge2(40)
	fmt.Println(p1.age) // 30  会值拷贝: 修改操作只是针对副本，无法修改接收者变量本身
}
