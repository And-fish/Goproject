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
// func belongs_to() {
// 	/*
// 		belongs to会与另一个模型建立一对一的连接。这种模型的每一个实例都属于另一个模型的实例
// 		例如，一个应用中包含了user和company，并且每个user能且只能分给一个company。在声明user结构体的时候要加上companyID的外键
// 		belongs to是属于的意思，他的外键在从表里，给user添加company的时候自动把company的id关联到companyid里
// */
// 	type Company struct {
// 		ID   int
// 		Name string
// 	}
// 	type User struct {
// 		gorm.Model
// 		Name      string
// 		CompanyID int
// 		Company   Company
// 	}
// }
// func belongs_to2() {
// 	/*
// 		定义一个belongs to的关系，数据库的表中必须存在外键。默认情况下，外键的名字使用拥有者的类型名加上表的主键的字段名字。
// 			也可以使用标签  foreignkey  自定义外键名字
// 		一般情况下主表的主键值作为外键的参考
// 			使用标签  references  来改变引用
// 		还可以使用OnUpdate和OnDelete来配置关联关系的级联操作
// 	*/
// 	type Company struct {
// 		ID   int
// 		Name string
// 	}
// 	type User struct {
// 		gorm.Model
// 		Name         string
// 		CompanyRefer int
// 		Company      Company `gorm:"foreignkey:CompanyRefer","references:Name","constraint":OnUpdate:CASCADE,OnDelete:SET NULL;`
// 	}
// }
