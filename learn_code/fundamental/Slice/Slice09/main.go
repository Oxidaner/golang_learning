package main

import (
	"fmt"
)

/*
	切片拷贝
*/
func main() {

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3 : %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3 : %v\n", s3)

	/*
		copy函数
	*/
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array data : ", data)
	s4 := data[8:]
	s5 := data[:5]
	fmt.Printf("slice s4 : %v\n", s4)
	fmt.Printf("slice s5 : %v\n", s5)
	copy(s5, s4)
	fmt.Printf("copied slice s4 : %v\n", s4)
	fmt.Printf("copied slice s5 : %v\n", s5)
	fmt.Println("last array data : ", data)
}
