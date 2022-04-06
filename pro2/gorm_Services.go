// package main

// import (
// 	"fmt"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// type tem struct {
// 	Name string
// }

// func init() {
// 	var err error
// 	// "用户名:密码@tcp(127.0.0.1:3306)/库名字?charset=utf8mb4&parseTime=True&loc=Local"
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/tset_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("GORM初始化成功,正在初始化表")
// 	}
// }
// func main() {
// 	/*
// 		gorm会在事务里执行写入操作(创建、更新、删除)
// 	*/

// 	//禁用默认事务，如果没有某个方面的要求，可以在初始化的时候禁用它，可以获得性能上的提升
// 	db.Create(&tem{Name: "try"}) //如果没有事务控制，会只插入这一条信息
// 	db.Create(nil)               //如果有事务控制，会都不插入，因为其中有一条无法执行

// 	//手动控制
// 	tx := db.Begin()
// 	tx.Create(&tem{Name: "try"})
// 	err := tx.Create(nil).Error
// 	if err != nil {
// 		tx.Rollback()
// 	} else {
// 		tx.Commit()
// 	}

// 	//自动控制
// 	db.Transaction(func(tx *gorm.DB) error {
// 		if err2 := tx.Create(&tem{Name: "try"}).Error; err2 != nil {
// 			return err2
// 		}
// 		if err2 := tx.Create(nil).Error; err2 != nil {
// 			return err2
// 		}
// 		return nil
// 	})

// 	//手动控制，savepoint设置回滚点,rollbackto设置回滚目标
// 	tx1 := db.Begin()
// 	tx.Create(&tem{Name: "try"})
// 	tx1.SavePoint("sp1")
// 	err = tx1.Create(nil).Error
// 	if err != nil {
// 		tx.RollbackTo("sp1")
// 	} else {
// 		tx.Commit()
// 	}
// }
