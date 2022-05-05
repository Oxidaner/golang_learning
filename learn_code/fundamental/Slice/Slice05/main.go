package main

import "fmt"

func main() {

	/*
		数组指针
	*/
	s := []int{0, 1, 2, 3}
	p := &s[2] // *int, 获取底层数组元素指针。
	*p += 100

	fmt.Println(s)

	/*
		至于 [][]T，是指元素类型为 []T 。
	*/
	data := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data)
}
