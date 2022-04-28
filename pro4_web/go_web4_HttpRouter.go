// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/julienschmidt/httprouter"
// )

// func test0() {
// 	r := httprouter.New()
// 	r.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		fmt.Fprint(w, "welcome")
// 	})
// 	r.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		fmt.Fprintf(w, "hello,%v", p.ByName("name"))
// 	})

// 	http.ListenAndServe(":8000", r)
// }
// func test1() {
// 	Index := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		fmt.Fprint(w, "Welcome\n")
// 	}
// 	Hello := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		fmt.Fprintf(w, "hello,%v\n", p.ByName("name"))
// 	}
// 	Getuser := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		fmt.Fprintf(w, "you are get user %v", p.ByName("uid"))
// 	}
// 	Adduser := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		fmt.Fprintf(w, "you are add user %v", p.ByName("uid"))
// 	}
// 	Deletuser := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		fmt.Fprintf(w, "you are delet user %v", p.ByName("uid"))
// 	}
// 	Modifyuser := func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		fmt.Fprintf(w, "you are modify user %v", p.ByName("uid"))
// 	}

// 	router := httprouter.New()
// 	router.GET("/", Index)
// 	router.GET("/hello/:name", Hello)

// 	router.GET("/getuser/:uid", Getuser)
// 	router.POST("/adduser/:uid", Adduser)
// 	router.DELETE("/deletuser/:uid", Deletuser)
// 	router.PUT("/modifyuser/:uid", Modifyuser)

// 	http.ListenAndServe(":2002", router)
// }
// func main() {
// 	/*
// 		HttpRouter是一种轻量级高性能的golangHTTP请求路由器。
// 		与Golang默认路由器相比，此路由器支持的路由模式中的变量并匹配请求方法。它还可以更好的扩展。
// 		该路由器针对高性能和小内存占用进行了优化，即使有很长的路径和大量的卢锡安，他也能很好地扩展。压缩动态特里(基数树)结构用于有效匹配
// 		gin框架就是以httprouter为基础开发的

// 		go get github.com/julienschmidt/httprouter
// 	*/
// 	test1()
// }
