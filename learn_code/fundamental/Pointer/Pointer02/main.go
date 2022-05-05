package main

import "fmt"

func modify1(x int) {
	//x = 100
	x++
}

func modify2(x *int) {
	*x = 100
}

func modify3(x int) int {
	x = 300
	return x
}

func main() {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}
