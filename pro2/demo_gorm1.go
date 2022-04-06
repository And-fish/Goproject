// package main

// import (
// 	"fmt"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type Product struct {
// 	gorm.Model
// 	/* 	type Model struct {
// 		ID        uint `gorm:"primarykey"`
// 		CreatedAt time.Time
// 		UpdatedAt time.Time
// 		DeletedAt DeletedAt `gorm:"index"`
// 	} */
// 	Code  string
// 	Price uint
// }

// var d *gorm.DB

// func gorm_init() {
// 	var err error
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Printf("初始化失败")
// 	}
// }
// func creation_table() {
// 	// 创建表
// 	d.AutoMigrate(&Product{})
// }
// func gorm_insert() {
// 	// 插入数据
// 	p1 := Product{
// 		Code:  "99999",
// 		Price: 999,
// 	}
// 	d.Create(&p1)
// }
// func gorm_find() {
// 	var p Product
// 	d.First(&p, 1) //根据主键查询，gorm.Model中的ID
// 	fmt.Printf("p: %v\n", p)
// 	p = Product{}
// 	d.First(&p, "code=?", "10086")
// 	d.Last(&p)
// 	fmt.Printf("p: %v\n", p)
// 	p = Product{}
// 	d.Where("code=?", "99999").Select("price", "code").Take(&p)
// 	fmt.Printf("p: %v\n", p)
// 	p1 := Product{Code: "10086"}
// 	d.Model(Product{}).Where(&p1).Take(&p) //根据已有的结构体用where条件查询
// 	fmt.Printf("p: %v\n", p)
// 	var c int64
// 	d.Model(Product{}).Count(&c)
// 	fmt.Printf("c: %v\n", c)

// 	// Select:筛选字段，只有筛选的字段是有赋值，其他都是空值
// 	// where:条件查询
// 	// take:设定结构体
// 	// Model:指定模型
// 	//	// Model(&struct{})和Model(&实例)的区别:里面是类型,则匹配类型,否则必须要完全一样
// 	// Count:计数
// }
// func gorm_update() {
// 	var ps []Product
// 	d.Limit(5).Find(&ps) //指定最大要查询的最大记录数
// 	for _, v := range ps {
// 		fmt.Printf("v: %v\n", v)
// 	}

// 	p := Product{}
// 	d.Model(&p).Where("code=?", "99999").Take(&p)
// 	d.Model(&p).Update("price", 999)
// 	fmt.Printf("p.Price: %v\n", p.Price)
// 	d.Model(&p).Updates(Product{Price: 555, Code: "666666"})
// 	fmt.Printf("p: %v\n", p)
// 	d.Model(&p).Updates(map[string]interface{}{"Price": 999, "code": "99999"})
// 	fmt.Printf("p: %v\n", p)
// 	p.Price = 200
// 	d.Save(&p)
// 	p2 := Product{Code: "55"}
// 	d.Save(&p2)
// 	fmt.Println("###########################################################")

// 	d.Limit(5).Find(&ps) //指定最大要查询的最大记录数
// 	for _, v := range ps {
// 		fmt.Printf("v: %v\n", v)
// 	}

// 	// updata(key,value)
// 	// updates(struct)
// 	// updates(map[string]interface{}{key:value})
// 	// save() 直接保存,之前不存在就插入,存在就更新
// }
// func gorm_del() {
// 	p := Product{}
// 	d.First(&p, "code=?", "55")
// 	d.Delete(&p) //软删除
// }
// func main() {
// 	/*
// 		ORM:对象关系映射简称ORM
// 			是一种为了解决面向对象与关系数据库(如mysql数据库)存在的互补匹配的现象的技术。
// 			简单来说,ORM十通过使用描述对象和数据库之间的映射的元数据,将持续中的对象自动持久化到关系数据库

// 		mysql语句:
// 			use go_db;
// 			drop table xxx;
// 			show tables;
// 			desc xxx;
// 			select * from xxx;
// 	*/

// 	gorm_init()
// 	gorm_del()

// }
