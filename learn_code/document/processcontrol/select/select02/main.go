package main

import (
	"fmt"
	"time"
)

//1. 超时判断
//比如在下面的场景中，使用全局resChan来接受response，如果时间超过3S,resChan中还没有数据返回，则第二条case将执行
var resChan = make(chan int)

// do request
func test() {
	select {
	case data := <-resChan:
		doData(data)
	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}
}

func doData(data int) {
	//...
}
