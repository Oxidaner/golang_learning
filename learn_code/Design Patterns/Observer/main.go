package main

import "fmt"

type PurchaseOperFunc func(status string, data string) (res bool, err error)

/**
 * 注册的观察者
 */
var PurchaseOperFuncArr = []PurchaseOperFunc{
	create,
	isDeleted,
	apply,
}

/**
创建的观察者
*/
func create(status string, data string) (res bool, err error) {
	if status == "create" {
		fmt.Println("开始创建")
		return true, nil
	}
	return true, nil
}

/**
 * 删除的观察者
 */
func isDeleted(status string, data string) (res bool, err error) {
	if status == "delete" {
		fmt.Println("开始删除")
		return true, nil
	}
	return true, nil
}

/**
 * 履约的观察者
 */
func apply(status string, data string) (res bool, err error) {
	if status == "apply" {
		fmt.Println("开始履约")
		return true, nil
	}
	return true, nil
}

func main() {
	status := "create"
	data := "订单数据"
	//有状态更新时，通知所有观察者
	for _, oper := range PurchaseOperFuncArr {
		res, err := oper(status, data)
		if err != nil {
			fmt.Println("操作失败")
			break
		}
		if res == false {
			fmt.Println("处理失败")
			break
		}
	}
}
