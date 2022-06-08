package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	ID        uint   // 字段名是 `id`
	Name      string // 字段名是 `name`
	Age       int
	Birthday  time.Time // 字段名是 `birthday`
	CreatedAt time.Time // 字段名是 `created_at`
}

func main() {
	db, err := gorm.Open("mysql", "zzb:123456@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var user = new(User)
	// 获取第一条记录，按主键排序
	db.First(user)
	fmt.Printf("%#v\n", user)
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，不指定排序
	db.Take(&user)
	fmt.Printf("user:%v\n", user)
	//// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录，按主键排序
	//db.Last(&user)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// 获取所有的记录
	//db.Find(&user)
	//// SELECT * FROM users;

	// 通过主键进行查询 (仅适用于主键是数字类型)
	//db.First(&user, 10)
	//// SELECT * FROM users WHERE id = 10;
}
