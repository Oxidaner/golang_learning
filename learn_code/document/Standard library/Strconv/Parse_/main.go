package main

import (
	"fmt"
	"strconv"
)

/*
	ParseBool()
	返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	func Parse_(str string) (value bool, err error)
*/

/*
	ParseInt()
	返回字符串表示的整数值，接受正负号。
*/

/*
	ParseUnit()
	ParseUint类似ParseInt但不接受正负号，用于无符号整型。
*/

/*
	ParseFloat()
	解析一个表示浮点数的字符串并返回其值。
*/
func main() {
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(b)
	}
	f, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(f)
	}
	i, err := strconv.ParseInt("-2", 10, 64)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(i)
	}
	u, err := strconv.ParseUint("2", 10, 64)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(u)
	}
}
