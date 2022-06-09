package main

import (
	"fmt"
	"io/ioutil"
)

func wr() {
	err := ioutil.WriteFile("./test.txt", []byte("www.5lmh.com"), 0666) //perm : 创建权限
	if err != nil {
		fmt.Println(err)
		return
	}
}

func re() {
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	wr()
	re()
}
