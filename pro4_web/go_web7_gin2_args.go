// package main

// import "github.com/gin-gonic/gin"

// func testGet(ctx *gin.Context) {
// 	// http://127.0.0.1:2002/testGet
// 	// username:,password:123456

// 	// http://127.0.0.1:2002/testGet?username=%E7%86%8A%E7%8E%BA%E7%9A%93&password=8012
// 	// username:熊玺皓,password:8012
// 	name := ctx.Query("username")
// 	pwd := ctx.DefaultQuery("password", "123456")
// 	ctx.String(200, "username:%v,password:%v", name, pwd)
// }
// func testPost(ctx *gin.Context) {
// 	// http://127.0.0.1:2002/testPost
// 	// Body[key:"userword",value:"熊玺皓"]
// 	// username:熊玺皓,password:123
// 	name := ctx.PostForm("username")
// 	pwd := ctx.DefaultPostForm("password", "123")
// 	ctx.String(200, "username:%v,password:%v", name, pwd)
// }
// func testPathParam(ctx *gin.Context) {
// 	// http://127.0.0.1:2002/hello/xxh/20
// 	// name:xxh,age:20
// 	name := ctx.Param("username")
// 	age := ctx.Param("age")
// 	ctx.String(200, "name:%v,age:%v", name, age)
// }
// func main() {
// 	e := gin.Default()
// 	e.LoadHTMLGlob("./gin_tmp/*")
// 	e.GET("/testGet", testGet)
// 	e.GET("/hello/:username/:age", testPathParam)
// 	e.POST("/testPost", testPost)

// 	e.GET("/goSearch", goSearch)
// 	e.POST("/search", Search)

// 	e.Run(":2002")
// }

// func goSearch(ctx *gin.Context) {
// 	ctx.HTML(200, "query.html", nil)
// }
// func Search(ctx *gin.Context) {
// 	page := ctx.DefaultQuery("page", "0")
// 	Keys := ctx.PostForm("key")
// 	ctx.String(200, "page is %v.\nkey is %v", page, Keys)
// }
