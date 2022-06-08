package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model //惯例模型 包含字段 ID，CreatedAt， UpdatedAt， DeletedAt。 可以自己定义
	Code       string
	Price      uint
}

func main() {
	db, err := gorm.Open("mysql", "zzb:123456@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//自动检查 Product 结构是否变化，变化则进行迁移
	db.AutoMigrate(&Product{})

	// 增
	//db.Create(&Product{Code: "L1212", Price: 1000})

	// 查
	var product Product
	db.First(&product, 1)                   // 找到id为1的产品
	db.First(&product, "code = ?", "L1213") // 找出 code 为 l1212 的产品

	// 改 - 更新产品的价格为 2000
	db.Model(&product).Update("Price", 5000)

	// 删 - 删除产品
	//db.Delete(&product)
}
