package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。
	func Print(a ...interface{}) (n int, err error)
	func Printf(format string, a ...interface{}) (n int, err error)
	func Println(a ...interface{}) (n int, err error)
*/

//func main() {
//	fmt.Print("在终端打印该信息。")
//	name := "枯藤"
//	fmt.Printf("我是：%s\n", name)
//	fmt.Println("在终端打印单独一行显示")
//}

/*
Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
	func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
*/

//func main() {
//	// 向标准输出写入内容
//	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
//	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
//	if err != nil {
//		fmt.Println("打开文件出错，err:", err)
//		return
//	}
//	name := "枯藤"
//	// 向打开的文件句柄中写入内容
//	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
//}

/*
Sprint系列函数会把传入的数据生成并返回一个字符串。
	func Sprint(a ...interface{}) string
	func Sprintf(format string, a ...interface{}) string
	func Sprintln(a ...interface{}) string
*/

//func main() {
//	s1 := fmt.Sprint("枯藤")
//	name := "枯藤"
//	age := 18
//	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
//	s3 := fmt.Sprintln("枯藤")
//	fmt.Println(s1, s2, s3)
//}

/*
	Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
	func Errorf(format string, a ...interface{}) error
*/

//func main() {
//	err := fmt.Errorf("这是一个错误")
//}

//func main() {
//	fmt.Printf("%v\n", 100)
//	fmt.Printf("%v\n", false)
//	o := struct{ name string }{"枯藤"}
//	fmt.Printf("%v\n", o)
//	fmt.Printf("%#v\n", o)
//	fmt.Printf("%T\n", o)
//	fmt.Printf("100%%\n")
//
//	n := 65
//	fmt.Printf("%b\n", n)
//	fmt.Printf("%c\n", n)
//	fmt.Printf("%d\n", n)
//	fmt.Printf("%o\n", n)
//	fmt.Printf("%x\n", n)
//	fmt.Printf("%X\n", n)
//}

//func main() {
//	var (
//		name    string
//		age     int
//		married bool
//	)
//	fmt.Scan(&name, &age, &married)
//	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
//
//
//}

func main() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
