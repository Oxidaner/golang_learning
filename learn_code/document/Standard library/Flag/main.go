package main

import (
	"fmt"
	"os"
)

// os.Args demo
// 获取命令行参数
func main() {
	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}
