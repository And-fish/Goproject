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
// func MTM() {
// 	/*
// 		many to many会在两个model中添加一张连接表。
// 		many2many标签会创建一个连接表
// 		这里只在user里面关联了language，当然也可以再在language关联user，不过数据量会变得很大
// 		可以使用foreignKey、references、joinforeignKey、jionreferences标签来重写外键和引用
// 	*/
// 	type Language struct {
// 		gorm.Model
// 		Name string
// 	}
// 	type User struct {
// 		gorm.Model
// 		Languages []Language `gorm:"many2many:user_languages;"`
// 	}
// }
// func MTM2() {
// 	type Language struct {
// 		gorm.Model
// 		Name string
// 	}
// 	type User struct {
// 		gorm.Model
// 		Languages []Language `gorm:"many2many:user_languages;"`
// 	}
// 	type UserLanguage struct {
// 		UserID     int `gorm:"primaryKey"`
// 		LanguageID int `gorm:"primaryKey"`
// 	}
// 	// 修改user的languages字段的连接表为userlanguage

// 	db.SetupJoinTable(&User{}, "languages", &UserLanguage{})
// }
