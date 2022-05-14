package main

import "fmt"

//人
type Person struct {
	name string
	sex  string
	age  int
}

// 自定义类型
type mystr string

// 学生
type Student struct {
	Person
	int
	mystr
}

func main() {
	s1 := Student{Person{"5lmh", "man", 18}, 1, "bj"}
	fmt.Println(s1)
}
