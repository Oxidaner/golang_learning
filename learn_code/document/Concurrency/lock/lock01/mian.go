package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}
func main() {
	wg.Add(2) // 增加2个goroutine
	go add()  // 启动一个goroutine
	go add()
	wg.Wait()
	fmt.Println(x)
}
