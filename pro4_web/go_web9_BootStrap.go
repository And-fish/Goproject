// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	/*
// 		https://getbootstrap.com/docs/5.1/getting-started/download/
// 		把下载的两个文件夹(js和css)放到当前文件夹下面的assets文件夹
// 	*/
// 	e := gin.Default()
// 	e.Static("/assets", "./assets")
// 	e.LoadHTMLFiles("./test_static.html")
// 	e.GET("go", func(ctx *gin.Context) {
// 		ctx.HTML(200, "test_static.html", nil)
// 	})
// 	e.Run(":2002")
// }
