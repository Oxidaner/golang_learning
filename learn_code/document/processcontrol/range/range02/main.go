package main

import "fmt"

func main() {
	a := [3]int{0, 1, 2}

	for i, v := range a { // index、value 都是从复制品中取出。

		if i == 0 { // 在修改前，我们先修改原数组。
			a[1], a[2] = 999, 999
			fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
		}

		a[i] = v + 100 // 使用复制品中取出的 value 修改原数组。

	}

	fmt.Println(a) // 输出 [100, 101, 102]。
}
