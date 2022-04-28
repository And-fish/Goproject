package router

import (
	"go_blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/assets", "./assets")

	// engine.GET("/", controller.Index)
	// engine.GET("/listuser", controller.ListUser)

	engine.GET("/register", controller.GoRegister)
	engine.POST("/register", controller.RegisterUser)

	engine.GET("/login", controller.GoLogin)
	engine.POST("/login", controller.Login)

	engine.GET("/", controller.Index)

	engine.GET("/post_index", controller.GetPostIndex)

	engine.GET("/post", controller.GoAddPost)
	engine.POST("/post", controller.AddPost)

	engine.GET("/post_detail", controller.PostDetail)
	engine.Run(":2002")
}
