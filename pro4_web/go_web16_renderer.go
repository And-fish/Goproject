// package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	/*
// 		Gin支持很多中华渲染，可以是简单的字符串、JSON、XML、HTML、protobuf。
// 			c.JSON(200,nil)
// 			c.XML(200,nil)
// 			c.HTML(200,"htmlfile.html",nil)
// 			c.String(200,"%v",nil)
// 			c.ProtoBuf(200,nil)
// 	*/
// 	e := gin.Default()
// 	e.GET("/render", func(c *gin.Context) {
// 		c.JSON(200, nil)
// 		c.XML(200, nil)
// 		c.HTML(200, "htmlfile.html", nil)
// 		c.String(200, "%v", nil)
// 		c.ProtoBuf(200, nil)
// 	})
// }
