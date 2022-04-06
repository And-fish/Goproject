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
// 	/*
// 		个人觉得，has one是把外键放在从表中，关联对象放在主表中，然后在从表结构声明的时候调用，可以把主表的某些字段的值同步到从表中
// 		belongs to是把外键和关联对象大批放在从表，然后在从表调用，把主表的某些字段同步到从表中
// 	*/
// }
// func Has_One() {
// 	/*
// 		has one与另一个模型建立一对一的关联，但是他和一对一关系不同，这种关联一个模型的每个实例都包含或者拥有另一个模型的实例
// 		has one是拥有的意思，外键在从表里，给user添加creditcard的时候，会把user的id关联到creditcard的userid里
// 	*/
// 	type CreditCard struct {
// 		gorm.Model
// 		Money  string
// 		UserID uint
// 	}
// 	type User struct {
// 		gorm.Model
// 		CreditCard CreditCard
// 	}
// }
// func has_one2() {
// 	/*
// 		定义一个belongs to的关系，数据库的表中必须存在外键。默认情况下，外键的名字使用拥有者的类型名加上表的主键的字段名字。
// 			也可以使用标签  foreignkey  自定义外键名字
// 		一般情况下主表的主键值作为外键的参考
// 			使用标签  references  来改变引用
// 		还可以使用OnUpdate和OnDelete来配置外键约束
// 	*/
// 	type CreditCard struct {
// 		gorm.Model
// 		Money     string
// 		UserRefer string
// 	}
// 	type User struct {
// 		gorm.Model
// 		CreditCard CreditCard `gorm:"foreignkey:UserRefer","references:Name","constraint":OnUpdate:CASCADE,OnDelete:SET NULL;`
// 	}
// }

// func has_one3() {
// 	/*
// 		gorm提供了多态关联支持，他会将游泳者的表名、主键值都保存到多态类型的字段中
// 		ploymorphicValue标签指定值，这时候添加cat的时候owner就不是cat而是master
// 	*/
// 	type Toy struct {
// 		ID        int
// 		Name      string
// 		OwnerID   int
// 		OwnerType string
// 	}
// 	type Cat struct {
// 		ID   int
// 		Name string
// 		Toy  Toy `gorm:"polymorphic:Owner;ploymorphicValue:master"`
// 	}
// 	type Dog struct {
// 		ID   int
// 		Name string
// 		Toy  Toy `gorm:"polymorphic:Owner;"`
// 	}
// }
