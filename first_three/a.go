package main

import (
	"first_three/mypath"
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func main() {
	a := 4
	fmt.Println("Hello, world", a)
	b := 7
	c := mypath.Add(a, b)
	fmt.Println(c)

	arr := []float64{1, 2, 4, 5, 6, 7, 8, 9, 10}
	v := stat.Variance(arr, nil)
	fmt.Printf("方差=%f\n", v)

	// 	要运行这个程序，先将将代码放到名为 hello-world.go 的文件中，然后执行 go run。
	// $ go run hello-world.go

	// 如果我们想将程序编译成二进制文件（Windows 平台是 .exe 可执行文件）， 可以通过 go build 来达到目的。
	// $ go build hello-world.go
	// $ ls

	// 然后我们可以直接运行这个二进制文件。
	// $ ./hello-world

}
