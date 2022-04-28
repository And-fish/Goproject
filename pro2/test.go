package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var d *gorm.DB

type TestObject struct {
	gorm.Model
	Name string
	Age  int
}

var default_object = TestObject{
	Name: "default",
	Age:  0,
}

func t(a ...string) {
	o1 := TestObject{
		Name: "1",
		Age:  6,
	}
	d.Select(a).Create(&o1)
}
func (u *TestObject) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("插入成功")
	return
}
func main() {
	var err error
	// "用户名:密码@tcp(127.0.0.1:3306)/库名字?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("GORW初始化成功,正在创建初始化表")
		d.AutoMigrate(&animal{})
	}

	// hashmap := map[string]interface{}{"name": "hello"}
	// keys := make([]string, 0, len(hashmap))
	// values := make([]interface{}, 0)
	// for k, v := range hashmap {
	// 	keys = append(keys, k)
	// 	values = append(values, v)
	// }
	// fmt.Printf("keys: %v\n", keys)
	// fmt.Printf("values: %v\n", values)

	// d.Select(keys).Create(&default_object)

	// 	ts := []TestObject{TestObject{Name: "1"}, TestObject{Name: "2"}}
	// 	d.Create(&ts)

	// ts := []TestObject{TestObject{Name: "6"}}
	// for i, v := range ts {
	// 	isexist := TestObject{}
	// 	d.Where("name=? and age=?", v.Name, v.Age).Take(&isexist)
	// 	if isexist.Name == v.Name {
	// 		fmt.Printf("数据%v已存在", i+1)
	// 		log.Fatal("gg")
	// 	}
	// 	fmt.Println("1")
	// 	d.Create(&v)
	// 	fmt.Println("2")
	// }
	// m := map[string]interface{}{"name": 7}
	// d.Model(&TestObject{}).Create(m)
	// d.Create(&TestObject{Name: "d"})

	// a := animal{Name: "dog", id: 1}
	// a2 := animal{Name: "cat", id: 2}
	// d.Create(&a)
	// d.Create(&a2)

	// var to = TestObject{
	// 	Name: "991",
	// }
	// d.Create(&to)

	// d.Unscoped().Last(&to)
	// d.Where(&to).Take(&to)
	// d.Unscoped().Delete(&to)
	// fmt.Printf("to: %v\n", to)
	// d.Delete(&to)
	// fmt.Printf("to: %v\n", to)
	cat := animal{
		"cat",
		4,
	}
	d.Table("animals").Create(&cat)
	var a animal
	d.Table("animals").Where(&cat).First(&a)
	fmt.Printf("a: %v\n", a)
}

type animal struct {
	Name string
	Id   int
}
