package dao

import (
	"fmt"
	model2 "gmall/repository/db/model"
	"os"
)

// Migration 执行数据迁移
func Migration() {
	// 自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model2.User{},
			&model2.Address{},
			&model2.Product{},
			&model2.Category{},
			&model2.Order{})
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
