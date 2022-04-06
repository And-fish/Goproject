// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var db *sql.DB

// type User struct {
// 	id       int
// 	username string
// 	passwd   string
// }

// func sql_open(user, passwd, dbname string) {
// 	/*
// 		open函数可能只是严重其参数格式是否正确，实际并不创建于数据库的连接。
// 		如果要检查数据源的名称是否真实有效，应该调用Ping方法
// 	*/
// 	s := fmt.Sprintf("%v:%v@/@%v", user, passwd, dbname) //"user:passwd@/dbname"
// 	db, err := sql.Open("mysql", s)
// 	// func sql.Open(driverName string, dataSourceName string) (*sql.DB, error)
// 	if err != nil {
// 		panic(err)
// 	}
// 	db.SetConnMaxLifetime(time.Minute * 3) //最大连接时间
// 	db.SetMaxOpenConns(10)                 //最大连接数
// 	db.SetMaxIdleConns(10)                 //空闲连接数
// 	fmt.Printf("db: %v\n", db)
// }
// func sql_init() (err error) {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
// 	db, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		return err
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }

// func queryOneRow(i int) {
// 	var u User
// 	s := "select * from user_tbl where id = ?"
// 	err := db.QueryRow(s, i).Scan(&u.id, &u.passwd, &u.username)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("u: %v\n", u)
// 	}
// }
// func queryManyRow() {
// 	s := "select * from user_tbl"
// 	r, err := db.Query(s)
// 	defer r.Close()
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		for r.Next() { //接下来准备使用Scan方法读取的下一个结果行。
// 			// 成功时返回true，如果没有下一个结果行或在准备时发生错误，则返回false。
// 			var u User
// 			r.Scan(&u.id, &u.passwd, &u.username)
// 			fmt.Printf("u: %v\n", u)
// 		}
// 	}
// }
// func sql_insertData(username, passwd string) {
// 	sqlStr := "insert into user_tbl(username,passwd) values (?,?)"
// 	r, err := db.Exec(sqlStr, username, passwd)
// 	// func (*sql.DB).Exec(query string, args ...interface{}) (sql.Result, error)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		i, _ := r.LastInsertId()
// 		i2, _ := r.RowsAffected()
// 		fmt.Printf("最后一行id为%v,更新了%v行\n", i, i2)
// 	}
// }
// func sql_updata(id, username, passwd string) {
// 	s := "update user_tbl set username=?, passwd=? where id=?"
// 	r, err := db.Exec(s, username, passwd, id)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Println("修改成功")
// 		i, _ := r.RowsAffected()
// 		fmt.Printf("更新了%v行\n", i)
// 	}
// }
// func sql_delete(id interface{}) {
// 	s := "delete from user_tbl where id=?"
// 	r, err := db.Exec(s, id)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Println("删除成功")
// 		i, _ := r.RowsAffected()
// 		fmt.Printf("更新了%v行\n", i)
// 	}
// }
// func main() {
// 	err := sql_init()
// 	if err != nil {
// 		fmt.Println("初始化失败")
// 	} else {
// 		fmt.Println("初始化成功")
// 	}

// 	// insertData("xiongxihao99", "8012wbym")
// 	queryOneRow(4)
// 	sql_updata("4", "xiongxihao66", "8012wbym")
// 	queryManyRow()
// 	sql_delete(5)
// 	queryManyRow()
// 	sql_insertData("xiongxiaho99", "8012wbym")
// }
