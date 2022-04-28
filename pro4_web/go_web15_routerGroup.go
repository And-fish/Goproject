// package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	/*
// 		假设你的网站有多个模块:博客\教程\视频等等,每个模块有多个路由,这样就可以进行路由分组
// 		其实在前面学习BasicAuth中间件的时候就已经使用过了
// 	*/
// 	e := gin.Default()
// 	rg := e.Group("/A_class")
// 	rg2 := e.Group("/B_class")

// 	rg.GET("/hello", func(ctx *gin.Context) {
// 		ctx.String(200, "hello ACLASS")
// 	})
// 	rg2.GET("/hello", func(ctx *gin.Context) {
// 		ctx.String(200, "hello BCLASS")
// 	})
// 	e.Run(":2002")
// 	// http://127.0.0.1:2002/hello
// 	// http://127.0.0.1:2002/A_class/hello
// 	// hello ACLASS
// 	// http://127.0.0.1:2002/B_class/hello
// 	// hello BCLASS
// }
