// package main

// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	/*
// 		中间件就是在请求中间起到拦截作用的处理函数
// 		gin.Default()默认两个中间件：
// 			Default返回一个已经附加了Logger和Recovery中间件的Engine实例。
// 			1.Logger():用来处理日志
// 			2.Recovery():用来处理错误

// 		如果使用gin.New():
// 			New返回一个新的空引擎实例，不附带任何中间件。

// 		默认配置为:
// 			—RedirectTrailingSlash: true
// 			—RedirectFixedPath: false
// 			—handlemethodnotalallowed: false
// 			—ForwardedByClientIP: true
// 			—UseRawPath: false
// 			—UnescapePathValues: true

// 		使用*gin.Engine.Use()就可以自定义选择中间件，在use函数中处理各种情况的方法就是自定义中间件
// 	*/
// 	e := gin.Default()
// 	fmt.Printf("e.RedirectTrailingSlash: %v\n", e.RedirectTrailingSlash)
// 	fmt.Printf("e.RedirectFixedPath: %v\n", e.RedirectFixedPath)
// 	fmt.Printf("e.HandleMethodNotAllowed: %v\n", e.HandleMethodNotAllowed)
// 	fmt.Printf("e.ForwardedByClientIP: %v\n", e.ForwardedByClientIP)
// 	fmt.Printf("e.UseRawPath: %v\n", e.UseRawPath)
// 	fmt.Printf("e.UnescapePathValues: %v\n", e.UnescapePathValues)

// 	e.Use(func(ctx *gin.Context) {
// 		fmt.Println("sorry")
// 	})
// 	e.GET("/mid", func(ctx *gin.Context) {
// 		ctx.String(200, "hello world")
// 	})
// 	e.Run(":2002")

// }
