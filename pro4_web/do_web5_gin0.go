// package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	/*
// 		gin-gonic.com
// 		Gin是一个golang的微框架，基于httprouter，封装比较优雅，API友好，源码注释比较明显，具有快速灵活，容错方便等特点。
// 		特征：
// 			1.速度快：基于基数树的路由，内存占用少。没有反射。可预测的API性能
// 			2.中间件支持：传入的HTTP请求可以由中间件链和最终操作处理。例如logger、gzip等最后在DB中发布的一条消息
// 			3.Crash-free：Gin可以捕获HTTP请求期间发生的panic并恢复它。这样你的服务器将始终可用
// 			4.JSON验证：Ginky解析和验证请求的JSON。例如，检查所需值的存在
// 			5.路由分组：更好地组织路线。需要授权和不需要授权，不同的API版本...此外，组可以无限嵌套，而不会降低性能
// 			6.错误管理：Gin提供了一种方便的方法来收集HTTP请求期间发生的错误。最终，中间件可以将它们写入日志文件、数据库并通过网络发送它们
// 			7.内置渲染：Gin为JSON、XML、HTML渲染提供了一个易于使用的API
// 			8.可扩展：创建一个新的中间件非常简单，只需查看示例代码即可

// 		got get github.com/gin-gonic/gin
// 	*/
// 	e := gin.Default()
// 	e.GET("/hello", func(ctx *gin.Context) {
// 		// ctx.String(200, "hello %v\n", "xxh")
// 		ctx.JSON(200, gin.H{
// 			"name": "xxh",
// 		})
// 	})
// 	//默认是8080端口
// 	e.Run(":2002")
// }
