// package main

// import (
// 	"log"

// 	"github.com/gin-gonic/gin"
// )

// func upload(c *gin.Context) {
// 	file, _ := c.FormFile("testfile")
// 	log.Println(file.Filename)

// 	// 上传到项目的根目录上面去
// 	c.SaveUploadedFile(file, file.Filename)

// 	c.String(200, "%v , uploaded", file.Filename)
// }
// func goupload(c *gin.Context) {
// 	c.HTML(200, "h2uploadfile.html", nil)
// }
// func main() {
// 	/*
// 		在form表单处设定method="post" enctype="multipart/form-data"
// 		在input 设置type="file"
// 		详细见h2uploadfile.html
// 	*/
// 	e := gin.Default()
// 	e.LoadHTMLFiles("./h2uploadfile.html")
// 	e.POST("/upload", upload)
// 	e.GET("/upload", goupload)
// 	e.Run(":2002")
// }
