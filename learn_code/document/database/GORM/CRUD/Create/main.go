package main

import (
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

	db.AutoMigrate(&User{})

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	print(db.NewRecord(user)) // => 返回 `true` ，因为主键为空

	db.Create(&user) // 将设置 `CreatedAt` 为当前时间
	//
	print(db.NewRecord(user)) // => 在 `user` 之后创建返回 `false`
}
