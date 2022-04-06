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
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/tset_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("GORM初始化成功,正在初始化表")
// 	}
// }
// func main() {
// 	/* 	db.AutoMigrate(&Student{}, &Email{}, &Address{}, &Language{})
// 	   	xxh := Student{
// 	   		Name: "熊玺皓",
// 	   		BillingAddress: Address{
// 	   			Address1: "Billing Address - Address 1",
// 	   		},
// 	   		ShippingAddress: Address{
// 	   			Address1: "Shipping Address - Address 1",
// 	   		},
// 	   		Emails: []Email{
// 	   			Email{Email: "1626245483@qq.com"},
// 	   			Email{Email: "helloworld@gmail.com"},
// 	   		},
// 	   		Languages: []Language{
// 	   			Language{Name: "ZH"},
// 	   			Language{Name: "EN"},
// 	   		},
// 	   	}
// 	   	db.Create(&xxh) */

// 	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&Student{}) //更新所有关联数据
// 	db.Model(&Student{}).Association("languages").Find(&[]Language{})         //查找所有匹配的关联记录
// 	db.Model(&Student{}).Association("languages").Append(&[]Language{
// 		Language{Name: "JP"},
// 	}) //为many2many和has many添加关联，为has one 和belongs to替换关联
// 	db.Model(&Student{}).Association("languages").Replace(&[]Language{
// 		Language{Name: "JP"},
// 	}) //用一个新的关联替换当前的关联
// 	db.Model(&Student{}).Association("languages").Delete(&[]Language{
// 		Language{Name: "EN"},
// 	}) //删除关联，只会删除引用，不会从数据库中删除这些对象
// 	db.Model(&Student{}).Association("languages").Clear() //清空关联记录
// 	db.Model(&Student{}).Association("languages").Count() //计数关联记录

// 	db.Select("languages").Delete(&Student{})                    //删除student
// 	db.Select("languages").Delete(&Student{Name: "hello"})       //删除student，不删除student的language
// 	db.Select("languages").Where("id = ?", 1).Delete(&Student{}) //删除student的同时，也删除student的language，因为只有记录的主键不为空关联才会被删除

// }

// type Email struct {
// 	//从表，has one
// 	gorm.Model
// 	Email     string
// 	StudentID int // 外键
// }
// type Address struct {
// 	//从表，has one
// 	gorm.Model
// 	Address1  string
// 	StudentID int //外键
// }
// type Language struct {
// 	gorm.Model
// 	Name     string
// 	Students []*Student `gorm:"many2many:student_languages;"` //many to many
// }
// type Student struct {
// 	gorm.Model
// 	Name            string
// 	BillingAddress  Address
// 	ShippingAddress Address
// 	Emails          []Email
// 	Languages       []Language `gorm:"many2many:student_languages;"` //many to many
// }
