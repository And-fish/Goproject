// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	/*
// 		BasicAuth中间件是用来提供对网络资源的访问保护的
// 	*/
// 	var datas = gin.H{
// 		"a": gin.H{"name": "xxh", "age": 20},
// 		"b": gin.H{"name": "xzy", "age": 10},
// 	}

// 	e := gin.Default()
// 	// Accounts defines a key/value for user/pass list of authorized logins.
// 	// 帐户为授权登录的用户/通过列表定义一个键/值。
// 	g := e.Group("/admin", gin.BasicAuth(gin.Accounts{
// 		"a": "1234",
// 		"b": "5678",
// 	}))

// 	// 访问http://127.0.0.1:8080/admin/hello会触发
// 	g.GET("/hello", func(ctx *gin.Context) {
// 		//AuthUserKey是basicauth中用户凭据的cookie名称。
// 		user := ctx.MustGet(gin.AuthUserKey).(string)

// 		data, ok := datas[user]
// 		if ok {
// 			ctx.JSON(200, gin.H{
// 				"user": user,
// 				"data": data,
// 			})
// 		} else {
// 			ctx.JSON(200, gin.H{
// 				"user": user,
// 				"data": gin.H{
// 					"data": "NO DATA :(",
// 				},
// 			})
// 		}
// 	})
// 	e.Run()

// }
