// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func how2getCookie(c *gin.Context) {
// 	name := c.Query("username")
// 	// Cookie returns the named cookie provided in the request or ErrNoCookie if not found. And return the named cookie is unescaped.
// 	// If multiple cookies match the given name, only one cookie will be returned.
// 	// Cookie返回请求中提供的已命名Cookie，如果没有找到则返回ErrNoCookie。并返回命名的cookie是未转义的。
// 	// 如果多个cookie匹配给定的名称，则只返回一个cookie。
// 	s, err := c.Cookie("username")
// 	if err != nil {
// 		s = name
// 		// cookie的名称、cookie的值、存活时间、路径、域、是否只能用https来访问、是否能用js读到
// 		c.SetCookie("username", s, 60*60, "/", "localhost", false, true)
// 	}
// 	c.String(200, "测试cookie,%v", name)
// }

// func main() {
// 	/*
// 		cookie是服务器向客户端写的1一些数据，可以实现像自动登录等功能
// 	*/
// 	e := gin.Default()
// 	e.GET("/h2gc", how2getCookie)
// 	e.Run(":2001")
// }
