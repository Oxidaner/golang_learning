package main

import "fmt"

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func main() {
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

// 导致关键流程出现不可修复性错误的使用 panic，其他使用 error。
