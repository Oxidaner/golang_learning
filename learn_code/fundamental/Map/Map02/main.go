package main

import "fmt"

func main() {
	/*
		Map遍历
	*/
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	/*
		只遍历key
	*/
	scoreMap = make(map[string]int)
	scoreMap["张四"] = 90
	scoreMap["大明"] = 100
	scoreMap["王六"] = 60
	for k := range scoreMap {
		fmt.Println(k)
	}
}
