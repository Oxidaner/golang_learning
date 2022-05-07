package main

import "fmt"

//Person 结构体Person类型
type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"pprof.cn",
		18,
	}
	fmt.Printf("%#v\n", p1)        //main.Person{string:"pprof.cn", int:18}
	fmt.Println(p1.string, p1.int) //pprof.cn 18
}
