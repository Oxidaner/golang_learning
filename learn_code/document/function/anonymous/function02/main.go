package main

import "fmt"

func test(s string, n ...int) (string, int) {
	var x int
	for _, i := range n {
		x += i
	}

	return s, x
}

func main() {
	s := []int{1, 2, 3}
	res, str := test("sum: %d", s...) // slice... 展开slice
	fmt.Printf(res, str)
}
