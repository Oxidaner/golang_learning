package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	//for i, j := range stus {
	//	fmt.Print(i, j)
	//}
	for _, stu := range stus {
		//m[stu.name] = &stu
		m[stu.name] = &stu //这里遍历过程中&stu会随stu改变,遍历到最后,&stu为"博客" 所有的都为&stu
		fmt.Printf("%v \n", &stu)
		//fmt.Printf("%#v \n", stu)
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
