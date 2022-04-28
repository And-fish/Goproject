// package main

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// 	Age  string `json:"age"`
// }

// var users = make([]User, 0)

// /*
// 	strconv.Atoi()将字符型转化为int型
// */
// func init() {
// 	u1 := User{"1", "xxh", "20"}
// 	u2 := User{"2", "tjn", "21"}
// 	u3 := User{"3", "qxx", "21"}
// 	u4 := User{"4", "ghs", "22"}
// 	users = append(users, u1)
// 	users = append(users, u2)
// 	users = append(users, u3)
// 	users = append(users, u4)
// 	fmt.Printf("users: %v\n", users)
// }

// func find(id string) (int, *User, bool) {
// 	for i, u := range users {
// 		if u.ID == id {
// 			return i, &u, true
// 		}
// 	}
// 	return -1, nil, false
// }
// func findGET(c *gin.Context) {
// 	c.HTML(200, "find.html", nil)
// }
// func findPOST(c *gin.Context) {
// 	id := c.PostForm("id")
// 	isFound := false
// 	for _, v := range users {
// 		if id == v.ID {
// 			isFound = true
// 			c.String(200, "查找成功\nname:%v\nid:%v,age:%v", v.Name, v.ID, v.Age)
// 			break
// 		}
// 	}
// 	if isFound == false {
// 		c.String(200, "没有此数据")
// 	}
// }

// func inset(c *gin.Context) {
// 	id := c.PostForm("id")
// 	name := c.PostForm("name")
// 	age := c.PostForm("age")
// 	_, _, isFound := find(id)
// 	if isFound {
// 		c.String(200, "数据可能已经存在于库中了,请重试")
// 	} else {
// 		users = append(users, User{
// 			ID:   id,
// 			Name: name,
// 			Age:  age,
// 		})
// 		c.String(200, "插入数据成功")
// 	}
// }
// func delete(c *gin.Context) {
// 	id := c.PostForm("id")
// 	i, _, isFound := find(id)
// 	if isFound {
// 		users = append(users[:i], users[i+1:]...)
// 		c.String(200, "删除数据成功")
// 	} else {
// 		_ = i
// 		c.String(200, "数据可能已经不存在于库中了,请重试")
// 	}
// }
// func chanage(c *gin.Context) {
// 	id := c.PostForm("id")
// 	name := c.PostForm("name")
// 	age := c.PostForm("age")
// 	i, _, isFound := find(id)
// 	if isFound {
// 		users = append(users[:i], users[i+1:]...)
// 		users = append(users, User{
// 			ID:   id,
// 			Name: name,
// 			Age:  age,
// 		})
// 		c.String(200, "删除数据成功")
// 	} else {
// 		_ = i
// 		c.String(200, "数据可能已经不存在于库中了,请重试")
// 	}
// }
// func show(c *gin.Context) {
// 	c.JSON(200, users)
// }
// func home(c *gin.Context) {
// 	c.HTML(200, "db_test.html", nil)
// }
// func main() {
// 	e := gin.Default()

// 	e.LoadHTMLGlob("./db_test/*")
// 	e.GET("/home", home)
// 	e.GET("/find", findGET)
// 	e.POST("/chanage", chanage)
// 	e.POST("/find", findPOST)
// 	e.POST("/inset", inset)
// 	e.POST("/delete", delete)
// 	e.GET("/allData", show)
// 	e.Run(":2002")
// }
