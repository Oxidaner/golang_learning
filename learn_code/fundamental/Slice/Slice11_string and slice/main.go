package main

import (
	"fmt"
)

func main() {
	/*
		string底层就是一个byte的数组，因此，也可以进行切片操作。
	*/
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:]
	fmt.Println(s2)

	/*
		tring本身是不可变的，因此要改变string中字符。需要如下操作： 英文字符串：
	*/
	str = "Hello world"
	s := []byte(str) //中文字符需要用[]rune(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)

	/*
		含有中文字符串
	*/
	str = "你好，世界！hello world！"
	t := []rune(str)
	t[3] = '够'
	t[4] = '浪'
	t[12] = 'g'
	t = t[:14]
	str = string(t)
	fmt.Println(str)
}
