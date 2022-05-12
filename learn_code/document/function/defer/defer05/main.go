package main

import "fmt"

func foo() (i int) {

	i = 0
	defer func() {
		fmt.Println(i)
	}()

	return 2
}

func main() {
	foo()
}
