// package main

// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var d *gorm.DB

// type Human struct {
// 	gorm.Model
// 	Name string
// 	Age  int
// }

// func gorm_init() {
// 	var err error
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Printf("初始化失败")
// 	}
// }
// func main() {
// 	/*
// 		GORM倾向于约定而不是配置.
// 		默认情况下,GORM使用ID作为主键,使用结构体的蛇形复数作为表名,字段名的蛇形作为列名
// 		即,结构体大写开头加复数作为表名,属性名大写作为列名

// 		值得一提的是,声明struct时,属性名一定要是大写开头的,否则会出现匹配不上的问题
// 		type Model struct {
// 			ID        uint `gorm:"primarykey"`
// 			CreatedAt time.Time
// 			UpdatedAt time.Time
// 			DeletedAt DeletedAt `gorm:"index"`
// 		}

// 		可以通过 `gorm:"标签"` 控制字段的权限
// 	*/

// }
