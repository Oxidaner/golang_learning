package main

import "fmt"

func main() {

	array := []int{10, 20, 30, 40}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n, slice)

	slices := make([]byte, 3)
	m := copy(slices, "abcdef")
	fmt.Println(m, slices)

	/*
		由于 Value 是值拷贝的，并非引用传递，所以直接改 Value 是达不到更改原切片值的目的的，
		需要通过 &slice[index] 获取真实的地址。
	*/
	slice = []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("value = %d , value-addr = %x , slice-addr = %x\n", value, &value, &slice[index])
	}
}
