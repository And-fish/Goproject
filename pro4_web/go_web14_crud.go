// package main

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func find(id int) (int, *User) {
// 	for i, v := range users {
// 		if id == v.ID {
// 			return i, &v
// 		}
// 	}
// 	return -1, nil
// }
// func get(c *gin.Context) {
// 	uid := c.Param("uid")
// 	id, _ := strconv.Atoi(uid)
// 	i, u := find(id)
// 	if i == -1 {
// 		c.String(200, "无法在数据库中查找到信息")
// 	} else {
// 		c.JSON(200, *u)
// 	}
// }

// func post(c *gin.Context) {
// 	uid := c.Param("uid")
// 	uname := c.Param("uname")
// 	uage := c.Param("uage")
// 	id, _ := strconv.Atoi(uid)
// 	i, u := find(id)
// 	if i == -1 {
// 		c.String(200, "无法在数据库中查找到信息")
// 	} else {
// 		u.ID = id
// 		u.Name = uname
// 		u.Age, _ = strconv.Atoi(uage)
// 		c.String(200, "修改成功")
// 	}
// }
// func put(c *gin.Context) {
// 	uid := c.Param("uid")
// 	uname := c.Param("uname")
// 	uage := c.Param("uage")
// 	id, _ := strconv.Atoi(uid)
// 	i, _ := find(id)
// 	if i == -1 {
// 		age, _ := strconv.Atoi(uage)
// 		users = append(users, User{
// 			ID:   id,
// 			Name: uname,
// 			Age:  age,
// 		})
// 		c.String(200, "添加成功")
// 	} else {
// 		c.String(200, "该数据可能已经存在")
// 	}
// }
// func delete(c *gin.Context) {
// 	uid := c.Param("uid")
// 	id, _ := strconv.Atoi(uid)
// 	i, _ := find(id)
// 	if i == -1 {
// 		c.String(200, "无法在数据库中查找到信息")
// 	} else {
// 		users = append(users[:i], users[i+1:]...)
// 	}
// }
// func main() {
// 	e := gin.Default()
// 	e.GET("/:uid/:name/:age", get)
// 	e.POST("/:uid/:name/:age", post)
// 	e.PUT("/:uid/:name/:age", put)
// 	e.DELETE("/:uid/:name/:age", delete)
// 	e.Run(":2002")
// }

// type User struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// var users = make([]User, 0)

// /*
// 	strconv.Atoi()将字符型转化为int型
// */
// func init() {
// 	u1 := User{1, "xxh", 20}
// 	u2 := User{2, "tjn", 21}
// 	u3 := User{3, "qxx", 21}
// 	u4 := User{4, "ghs", 22}
// 	users = append(users, u1)
// 	users = append(users, u2)
// 	users = append(users, u3)
// 	users = append(users, u4)
// 	fmt.Printf("users: %v\n", users)
// }
