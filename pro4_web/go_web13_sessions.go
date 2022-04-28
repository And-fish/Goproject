// package main

// import (
// 	"github.com/gin-contrib/sessions"
// 	"github.com/gin-contrib/sessions/cookie"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	/*
// 		http是无状态、短连接，可以使用session保存客户端和服务器直接的会话状态
// 	*/
// 	e := gin.Default()
// 	store := cookie.NewStore([]byte("secret"))
// 	e.Use(sessions.Sessions("MySession", store))

// 	e.GET("/hello", func(ctx *gin.Context) {
// 		session := sessions.Default(ctx)
// 		// 用get取session信息
// 		if session.Get("hello") != "world" {
// 			// 设置session内容
// 			session.Set("hello", "world")
// 			// 保存
// 			session.Save()
// 		}
// 		ctx.JSON(200, gin.H{
// 			"hello": session.Get("hello"),
// 		})
// 	})

// 	e.Run(":2002")
// }
