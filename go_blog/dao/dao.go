package dao

import (
	"go_blog/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 接口
type Manager interface {
	// 用户操作
	AddUser(user *model.User)
	Login(username string) model.User

	// 博客操作
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	Getpost(pid int) model.Post
}

// 结构体
type manager struct {
	db *gorm.DB
}

// 接口实例,调用方法的时候直接调用通过这个接口示例完成
var Mgr Manager

// 连接到数据库
func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/web_user_data?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

// 用户操作
// manager结构体实现接口
func (mgr *manager) AddUser(user *model.User) {
	mgr.db.Create(user)
}
func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

// 博客操作:
func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}
func (mgr *manager) GetAllPost() []model.Post {
	posts := make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}
func (mgr *manager) Getpost(pid int) model.Post {
	var post model.Post
	mgr.db.First(&post, pid)
	return post
}
