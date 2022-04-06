// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"math"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB
// var nilstring string
// var nilint int

// type User struct {
// 	gorm.Model
// 	Name string
// 	Age  int
// 	/* 	// birthday time.Time */
// 	Birthday time.Time
// }

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
// 		db.AutoMigrate(&User{})
// 	}
// }
// func retrieve() User {
// 	u := User{Name: "熊辉"}
// 	db.First(&u)
// 	// SELECT * FROM users ORDRY BY id LIMIT 1;

// 	u = User{}
// 	db.Take(&u)
// 	// SELECT * FROM users LIMIT 1;

// 	u = User{}
// 	db.Last(&u)
// 	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

// 	fmt.Printf("返回查找到的记录数: %v\n", db.First(&User{}).RowsAffected)
// 	fmt.Printf("return err or nil: %v\n", db.First(&User{}).Error)

// 	/*
// 		First和Last根据主键排序，只有在目标struct是指针或者通过db.Model()才能使用
// 	// */
// 	// u=User{}
// 	// us=[]User{}
// 	db.First(&User{}) // 可以使用，因为目标struct是指针

// 	m := map[string]interface{}{}
// 	db.Model(&User{}).First(&m) // 可以使用，因为使用了Model只当了model为user
// 	fmt.Printf("m: %v\n", m)
// 	//SELECT *FEOM `users` ORDER BY `users`.`id` LIMIT1

// 	m = map[string]interface{}{}
// 	db.Table("users").First(&m) // 不能使用，因为无法关联
// 	fmt.Printf("m: %v\n", m)

// 	m = map[string]interface{}{}
// 	db.Table("users").Take(&m) // 可以使用，take可以直接存入
// 	fmt.Printf("m: %v\n", m)

// 	u = User{}
// 	db.First(&u, 4) // 主键是int类型才行，否则会先查
// 	fmt.Printf("u: %v\n", u)
// 	// SELECT * FROM `users` WHERE id = 1 ORDER BY id LIMIT 1
// 	// SELECT * FROM `表名` WHERE 查询条件 ORDER BY `表名`.`主键名` LIMIT 1
// 	// 所以根据上面的语法规则，第一个表名是默认了的，最后一个表名+主键名也是默认了的。
// 	// 所以只要在first里把参数传到sql语句里，即使主键不是int类型也能用fiest直接插
// 	// 例如我想查表名为animals主键名为name的第一个记录，可以这样写: 	db.First(&animal{}, "`name`=\"dog\"")

// 	us := []User{}
// 	db.Find(&us, []int{1, 2, 3})
// 	for _, v := range us {
// 		fmt.Printf("v: %v\n", v)
// 	}
// 	u = User{}
// 	db.Find(&u, 5)
// 	// find可以查询多个也可以查询单个，也是根据主键查询

// 	//Where
// 	db.Where("name=?", "xxh").First(&User{})
// 	db.Where("name=?", "xxh").Find(&User{})
// 	db.Where("name IN ?", []string{"xxh", "熊子媛"}).Find(&[]User{})
// 	db.Where("name LIKE ?", "%子%").Find(&User{})
// 	us = []User{}
// 	db.Where("age>=? And name=?", "19", "xxh").Find(&us)
// 	for _, v := range us {
// 		fmt.Printf("v: %v\n", v)
// 	}
// 	us = []User{}
// 	db.Where("birthday> ?", "2000-08-20").Find(&us) //时间也可以
// 	for _, v := range us {
// 		fmt.Printf("v: %v\n", v)
// 	}
// 	us = []User{}
// 	db.Where("birthday Between ? and ?", "2002-03-20", "2004-01-01").Find(&us)
// 	for _, v := range us {
// 		fmt.Printf("v: %v\n", v)
// 	}

// 	u = User{}
// 	db.Where(&User{Name: "熊辉"}).First(&u)
// 	u = User{}
// 	db.Where(map[string]interface{}{"name": "熊子媛"}).First(&u)
// 	db.Where([]int{1, 2, 3}).Find(&[]User{}) //一切查询，如果不声明查讯的列名，都默认是主键
// 	// 如果查询的条件包含0值(int的0、string的""之类的)，都不能用结构体的方法查询，否则会被忽略，建议使用map
// 	u = User{}
// 	db.Where(User{Name: "xxh"}, "name", "age").First(&u) // 当然也可以用这种方法查询到age=0的记录
// 	fmt.Printf("u: %v\n", u)
// 	u = User{}
// 	db.Where(User{Name: "xxh"}, "age").First(&u) // 和上面的不同，没有使用name就不会使用name来查，即使在结构体里面声明了
// 	fmt.Printf("u: %v\n", u)

// 	//内联
// 	db.First(&User{}, "id=?", 3)
// 	db.Find(&[]User{}, "name <> ? AND age > ?", "xxh", 10)
// 	db.Find(&[]User{}, User{Name: "xxh"})
// 	db.Find(&[]User{}, map[string]interface{}{"age": 20})

// 	// NOT 和 OR
// 	db.Not("name = ?", "xxh").Find(&[]User{})
// 	db.Where("name = ?", "xxh").Or("name = ?", "熊子媛").Find(&[]User{})

// 	//Select	指定读取的字段，默认是所有
// 	db.Select("name", "age").Find(&[]User{})

// 	//Order
// 	db.Order("age desc,name").First(&User{}) // 按照年龄倒着排序，年龄相同按照名字排序

// 	//Limit和Offset	指定查询数量和指定跳过数量
// 	u1 := []User{}
// 	u2 := []User{}
// 	db.Limit(3).Find(&u1).Limit(-1).Find(&u2)
// 	for _, v := range u1 {
// 		fmt.Printf("v1: %v\n", v)
// 	}
// 	for _, v := range u2 {
// 		fmt.Printf("v2: %v\n", v)
// 	}
// 	u1 = []User{}
// 	u2 = []User{}
// 	db.Offset(3).Find(&u1).Offset(-1).Find(&u2)
// 	for _, v := range u1 {
// 		fmt.Printf("v1: %v\n", v)
// 	}
// 	for _, v := range u2 {
// 		fmt.Printf("v2: %v\n", v)
// 	}

// 	// Group和Having 去重
// 	var re struct {
// 		Name  string
// 		Total int
// 	}
// 	db.Table("test_objects").Select("name,sum(age) as total").Group("name").Having("total = ?", "62").First(&re)
// 	// 把test_objects表中name一样的按照name和sum(age)查找出来,把其中name为10的打印出来
// 	//where里的字段必须是table中存在的，haveing只能用select中的
// 	fmt.Printf("re: %v\n", re)

// 	//disyinct
// 	var v1 struct {
// 		Name string
// 		Age  int
// 	}
// 	db.Table("test_objects").Select("name", "age").Distinct("name", "age").Order("age desc").Find(&v1)
// 	//去除distinct中字段重复项，去除了name和age都相同的字段
// 	fmt.Printf("v1: %v\n", v1)

// 	//scan
// 	var res struct {
// 		Name string
// 		Age  int
// 	}
// 	db.Model(&u).Select("name", "age").Order("id").First(&res)
// 	fmt.Printf("res: %v\n", res)
// 	// 如果struct没有主键，会使用第一个字段排序
// 	return u
// }

// func try_insert(name, birthday string) {
// 	isexist := User{}
// 	db.Where("birthday=?", birthday).Where("name=?", name).Take(&isexist)
// 	if isexist.Name != nilstring {
// 		fmt.Println("数据已存在")
// 		return
// 	} else {
// 		loc, _ := time.LoadLocation("Asia/Shanghai")
// 		t_birthday, _ := time.ParseInLocation("2006-01-02", birthday, loc)
// 		age := time.Now().Sub(t_birthday).Hours()
// 		age = math.Floor(age / 24 / 365)
// 		i := int(age)
// 		u := User{
// 			Name:     name,
// 			Age:      i,
// 			Birthday: t_birthday,
// 		}
// 		db.Create(&u)
// 		fmt.Println("插入成功")
// 		return
// 	}
// }
// func inset_manyData(users ...User) {
// 	// db.Create(&users)
// 	// db.CreateInBatches(users,100)	一批插入100条
// 	for i, v := range users {
// 		isexist := User{}
// 		db.Where("name=? and birthday=?", v.Name, v.Birthday).Take(&isexist)
// 		if isexist.Name == v.Name {
// 			fmt.Printf("数据%v已存在\n", i)
// 		} else {
// 			db.Create(&v)
// 		}
// 	}
// }
// func inset_select(keys ...string) {
// 	loc, _ := time.LoadLocation("Asia/Shanghai")
// 	d_birthday, _ := time.ParseInLocation("2006-01-02", "0000-00-00", loc)
// 	default_user := User{
// 		Name:     "default",
// 		Age:      -1,
// 		Birthday: d_birthday,
// 	}
// 	db.Select(keys).Create(&default_user)
// }
// func inset_byModel(users map[string]interface{}) {
// 	// 不会自动填充主键(没声明的为null,会填充id)
// 	// 可以批量 	db.Model(&User{}).Create([]map[string]interface{})
// 	db.Model(&User{}).Create(users)
// }
// func update(id int, newvalues ...setContainer) {
// 	var u = User{}
// 	db.First(&u, id)
// 	var o = option{
// 		Name:     u.Name,
// 		Age:      u.Age,
// 		Birthday: u.Birthday,
// 	}
// 	for _, v := range newvalues {
// 		v.array(&o)
// 	}
// 	u.Name = o.Name
// 	u.Age = o.Age
// 	u.Birthday = o.Birthday
// 	db.Save(&u)

// 	/* 	d2 := db.Model(&User{}).Where("id>?", 1).Update("updated_at", time.Now())
// 	   	db.Model(&u).Update("updated_at", time.Now().Add(time.Minute))
// 	   	db.Model(&u).Updates(map[string]interface{}{"updated_at": time.Now().Add(time.Minute * 2)})
// 	   	db.Model(&u).Select("name").Updates(map[string]interface{}{"updated_at": time.Now().Add(time.Hour * 2)})
// 	   	db.Model(&u).Omit("updated_at").Updates(map[string]interface{}{"updated_at": time.Now().Add(time.Hour * 9)})
// 	   	// 如果没有加任何条件的情况下gorm是不允许批量更新的
// 	   	fmt.Printf("更新了%v条记录\n", d2.RowsAffected)

// 	   	db.Model(&u).Update("age", gorm.Expr("age*?+?", 2, 9))	// age*2+9
// 	   	db.Model(&u).UpdateColumn("name","熊玺皓")	// 不使用时间追踪更新
// 	   	db.Model(&[]User{}).Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"},{Name: "age"}}}).Where("id < ?",0).Update("age",gorm.Expr("age*?",5))
// 	   	// 返回修改的到[]User{}, */
// }

// func gorm_Delete(users ...User) {
// 	for _, u := range users {
// 		// 软删除
// 		db.Delete(&u)
// 		db.Delete(&User{}, u.ID)

// 		db.Delete(User{}, "name = ?", "test") // 批量删除，如果没有任何条件gorm不会执行该操作

// 		// 真正的删除，id会空出来
// 		db.Unscoped().Delete(&u)
// 	}
// 	var u User
// 	db.Unscoped().Last(&u) // Unscoped可以查找到已经被软删除的记录
// 	fmt.Printf("u: %v\n", u)
// }

// func raw_sql() {
// 	type Result struct {
// 		ID   int
// 		Name string
// 		Age  int
// 	}

// 	var r1 Result
// 	db.Raw("select * from users where id > ? limit 1", 4).Scan(&r1)
// 	fmt.Printf("r1: %v\n", r1)

// 	var a1 int
// 	db.Raw("select count(age) from users where age < 20 and age > 0").Scan(&a1)
// 	fmt.Printf("a1: %v\n", a1)

// 	var u1 User
// 	db.Where("name = @myname", sql.Named("myname", "熊玺皓")).Find(&u1)
// 	fmt.Printf("u1: %v\n", u1)

// 	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
// 		return tx.Model(&User{}).Where("id > ?", 1).Limit(2).Order("age desc").Find(&[]User{})
// 	})
// 	fmt.Printf("sql: %v\n", sql)

// 	var u2 User
// 	var name string
// 	var age int
// 	row := db.Table("users").Select("name", "age").Where("id > 2").Order("age desc").Find(&u2).Row()
// 	row.Scan(&name, &age)
// 	fmt.Printf("u2: %v\n", u2)
// 	fmt.Printf("name: %v\n", name)
// 	fmt.Printf("age: %v\n", age)

// 	r, _ := db.Model(&User{}).Where("age > ?", 18).Select("name", "age").Rows()
// 	defer r.Close()
// 	var u3 User
// 	for r.Next() {
// 		db.ScanRows(r, &u3)
// 		fmt.Printf("u3: %v\n", u3)
// 		r.Scan(&name, &age)
// 		fmt.Printf("name: %v\n", name)
// 		fmt.Printf("age: %v\n", age)
// 	}
// }

// func main() {
// 	try_insert("熊辉", "1977-09-08")
// 	// inset_select("name", "age", "brithday")
// 	// a := User{Name: "xxh"}
// 	// inset_manyData(a)
// 	China := get_NewUserinfo("China", "1949-10-01")
// 	inset_manyData(China)
// 	// test := get_NewUserinfo("test", "2077-01-01")
// 	// db.Create(&test)
// 	// fmt.Printf("retrieve(): %v\n", retrieve())
// 	update(1, setName("熊玺皓"))
// 	// gorm_Delete()

// 	// var u User
// 	// db.Unscoped().Last(&u)	// Unscoped可以查找到已经被软删除的记录
// 	// fmt.Printf("u: %v\n", u)

// 	raw_sql()
// }

// type option struct {
// 	Name     string
// 	Age      int
// 	Birthday time.Time
// }
// type setContainer interface {
// 	array(*option)
// }
// type optionsFunc func(*option)

// func (f optionsFunc) array(o *option) {
// 	f(o)
// }
// func setName(name string) setContainer {
// 	return optionsFunc(func(o *option) {
// 		o.Name = name
// 	})
// }
// func setBirthday(birthday string) setContainer {
// 	return optionsFunc(func(o *option) {
// 		loc, _ := time.LoadLocation("Asia/Shanghai")
// 		t_birthday, _ := time.ParseInLocation("2006-01-02", birthday, loc)
// 		age := time.Now().Sub(t_birthday).Hours()
// 		age = math.Floor(age / 24 / 365)
// 		i := int(age)
// 		o.Age = i
// 		o.Birthday = t_birthday
// 	})
// }
// func NewUserinfo(ops ...setContainer) User {
// 	loc, _ := time.LoadLocation("Asia/Shanghai")
// 	f_birthday, _ := time.ParseInLocation("2006-01-02", "0000-00-00", loc)
// 	default_option := option{
// 		Name: "UNknown",
// 		Age:  -1, //最大岁数为292，因为time.time在time包里是用time.Duration储存的，
// 		// time.Duration的纳秒表示int的1,所以最多能保存9223372036854775807，也就是292年
// 		Birthday: f_birthday,
// 	}
// 	for _, v := range ops {
// 		v.array(&default_option)
// 	}
// 	return User{
// 		Name:     default_option.Name,
// 		Age:      default_option.Age,
// 		Birthday: default_option.Birthday,
// 	}
// }
// func get_NewUserinfo(name, birthday string) User {
// 	new_u := NewUserinfo()
// 	switch {
// 	case name != nilstring && birthday != nilstring:
// 		n := setName(name)
// 		t := setBirthday(birthday)
// 		new_u = NewUserinfo(n, t)
// 	case name == nilstring && birthday != nilstring:
// 		t := setBirthday(birthday)
// 		new_u = NewUserinfo(t)
// 	case name != nilstring && birthday == nilstring:
// 		n := setName(name)
// 		new_u = NewUserinfo(n)
// 	default:
// 	}
// 	return new_u
// }
