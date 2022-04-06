// package main

// import (
// 	"fmt"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func init() {
// 	var err error
// 	// "用户名:密码@tcp(127.0.0.1:3306)/库名字?charset=utf8mb4&parseTime=True&loc=Local"
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("GORM初始化成功,正在初始化表")
// 		// db.AutoMigrate(&User{})
// 	}
// }
// func main() {

// }
// func has_many() {
// 	/*
// 		has many与另一个模型建立了一对多的连接。不同于has many，拥有者可以有零或多个关联模型。
// 	*/
// 	type CreditCard struct {
// 		ID     int
// 		Money  int
// 		UserID int
// 	}
// 	type User struct {
// 		gorm.Model
// 		CreditCard []CreditCard
// 	}
// }
