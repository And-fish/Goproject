// package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	e := gin.Default()
// 	e.LoadHTMLGlob("./gin_tmp/*")
// 	e.GET("/login", func(ctx *gin.Context) {
// 		ctx.HTML(200, "login.html", nil)
// 	})
// 	e.POST("/login", func(ctx *gin.Context) {
// 		username := ctx.PostForm("username")
// 		password := ctx.PostForm("password")
// 		ctx.HTML(200, "welcome.html", gin.H{
// 			"username": username,
// 			"passwoed": password,
// 		})
// 	})
// 	e.Run(":2002")
// }
