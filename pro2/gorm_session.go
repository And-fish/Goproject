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
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("GORM初始化成功,正在初始化表")
// 	}
// }
// func main() {
// 	/*
// 		gorm提供了Session方法，这是一个New Session Method，它允许创建带配置的新建会话模式
// 		type Session struct {
// 			DryRun                   bool		//只生成sql但不执行
// 					stmt := db.Session(&gorm.Session{DryRun: true}).First(&tem{}).Statement
// 					stmt.SQL.String()
// 			PrepareStmt              bool		//在执行任何sql时都会创建一个prepared statement并将其缓存，以提高后续的效率
// 					tx := db.Session(&gorm.Session{PrepareStmt: true})
// 					tx.First(&tem{})
// 			NewDB                    bool		//创建一个不带之前条件的新DB
// 					tx := db.Where("name = ?","xxh").Session(&gorm.Session{NewDB: true})
// 					tx.First(&tem{})	tx是不包含where查找条件的
// 			Initialized              bool
// 			SkipHooks                bool		//跳过hook
// 			SkipDefaultTransaction   bool		//禁用全局事务
// 			DisableNestedTransaction bool		//禁用嵌套事务
// 			AllowGlobalUpdate        bool		//允许全局更新
// 			FullSaveAssociations     bool		//更新关联数据
// 			QueryFields              bool
// 			Context                  context.Context		//传入context哎追踪sql操作
// 			Logger                   logger.Interface		//自定义内建logger
// 			NowFunc                  func() time.Time		//允许改变获取当前时间的实现
// 					db.Where("name = ?", "xxh").Session(&gorm.Session{NowFunc:func() time.Time {return time.Now().Local()}})

// 			CreateBatchSize          int		//一次最多查几个
// 		}
// 	*/
// }
