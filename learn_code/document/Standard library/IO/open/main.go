package main

import (
	"fmt"
	"os"
)

func main() {
	//var buf [16]byte
	////os.Stdin：标准输入的文件实例，类型为*File
	//os.Stdin.Read(buf[:])
	////os.Stdin.WriteString(string(buf[:]))
	//
	//fmt.Println(string(buf[:]))

	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("abc.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	file.Close()
}
