// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	Name     string   `form:"name"`
// 	Password string   `form:"pwd"`
// 	Hobby    []string `form:"hobby"`
// 	Gender   string   `form:"gender"`
// 	City     string   `form:"city"`
// }

// func gotest(c *gin.Context) {
// 	c.HTML(200, "test_from.html", nil)
// }
// func test(c *gin.Context) {
// 	name := c.PostForm("name")
// 	pwd := c.PostForm("pwd")
// 	hobby := c.PostFormArray("hobby")
// 	gender := c.PostForm("gender")
// 	city := c.PostForm("city")
// 	c.HTML(200, "test_show.html", gin.H{
// 		"name":   name,
// 		"pwd":    pwd,
// 		"hobby":  hobby,
// 		"gender": gender,
// 		"city":   city,
// 	})
// }
// func test2(c *gin.Context) {
// 	//绑定结构体
// 	var user User
// 	c.ShouldBind(&user)
// 	c.String(200, "User:%s", user)
// }
// func main() {
// 	e := gin.Default()
// 	e.LoadHTMLFiles("./test_from.html", "./test_show.html")
// 	e.GET("/gotest", gotest)
// 	e.POST("/test", test)
// 	e.POST("/test2", test2)

// 	e.POST("/testpost", testPostStruct)
// 	e.GET("/testget", testGetstruct)
// 	e.GET("/testget/:username/:city", testGetUri)
// 	e.Run(":2002")
// }
// func testPostStruct(c *gin.Context) {
// 	// http://127.0.0.1:2002/testpost
// 	// kty:name,value:熊玺皓
// 	// User:{熊玺皓  []  }
// 	var user User
// 	c.ShouldBind(&user)
// 	c.String(200, "User:%s", user)
// }
// func testGetstruct(c *gin.Context) {
// 	// http://127.0.0.1:2002/testget?name=hello&city=changsha
// 	/* 	username:hello
// 	usercity:changsha */
// 	var user User
// 	c.ShouldBind(&user)
// 	c.String(200, "username:%v\nusercity:%v", user.Name, user.City)
// }
// func testGetUri(c *gin.Context) {
// 	// http://127.0.0.1:2002/testget/xxh/shanghai
// 	/* 	username:xxh
// 	usercity:shanghai */
// 	type Ur struct {
// 		Name string `uri:"username"`
// 		City string `uri:"city"`
// 	}
// 	var ur Ur
// 	c.ShouldBindUri(&ur)
// 	c.String(200, "username:%v\nusercity:%v", ur.Name, ur.City)
// }
