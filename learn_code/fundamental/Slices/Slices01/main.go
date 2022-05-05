package main

import "fmt"

func main() {
	/*
		赋值和函数传参操作都会复制整个数组数据。
	*/
	arrayA := [2]int{100, 200}
	var arrayB [2]int

	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	testArray(arrayA)

	/*
		函数传参用数组的指针
	*/
	arrayC := []int{100, 200}
	testArrayPoint(&arrayC) // 1.传数组指针
	arrayD := arrayC[:]
	testArrayPoint(&arrayD) // 2.传切片
	fmt.Printf("arrayC : %p , %v\n", &arrayC, arrayC)
}

func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}

func testArrayPoint(x *[]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}
