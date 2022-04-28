/*
	这个包是用来处理数据的handle
*/
package controller

import (
	"fmt"
	"go_blog/dao"
	"go_blog/model"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

// 用户部分
// 注册用户
func RegisterUser(c *gin.Context) {
	// 获取数据
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Printf("username: %v\n", username)
	fmt.Printf("password: %v\n", password)
	newuser := model.User{
		Username: username,
		Password: password,
	}
	dao.Mgr.AddUser(&newuser)
	// 跳转页面
	c.Redirect(301, "/")
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

// 登录用户]
func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := dao.Mgr.Login(username)

	if user.Username == "" {
		c.HTML(200, "login.html", "用户不存在")
	} else {
		if user.Password != password {
			c.HTML(200, "login.html", "密码错误")
		} else {
			c.Redirect(301, "/")
		}
	}
}

// 首页
func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

//博客部分
func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postindex.html", posts)
}
func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}
func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	context := c.PostForm("context")

	Newpost := model.Post{
		Title:   title,
		Tag:     tag,
		Context: context,
	}
	dao.Mgr.AddPost(&Newpost)

	c.Redirect(302, "/post_index")
}
func PostDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	post := dao.Mgr.Getpost(pid)

	context := blackfriday.Run([]byte(post.Context))

	c.HTML(200, "postdetail.html", gin.H{
		"Title":   post.Title,
		"Context": template.HTML(context),
	})
}
