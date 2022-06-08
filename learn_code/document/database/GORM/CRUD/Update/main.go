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

	var user User
	db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)

	fmt.Printf("user:%#v\n", user)
}
