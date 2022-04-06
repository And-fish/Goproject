// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func log_print() {
// 	log.Print("my log")
// 	log.Printf("my log %v", 100)
// 	name := "dasd"
// 	age := 20
// 	log.Println(name+" is", age, "years old")
// }
// func log_panic() {
// 	// defer fmt.Println("defer")	//但是能够执行defer
// 	fmt.Println("hello")
// 	log.Panic("get panic")
// 	// 退出进程抛出panic，接下来的程序不会运行

// 	defer fmt.Println("end...")
// }
// func log_fatal() {
// 	defer fmt.Println("defer......")
// 	fmt.Println("hello")
// 	log.Fatal("fatal...........")
// 	// 没有执行defer
// 	fmt.Println("end")
// }
// func log_SetFlags() {

// 	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
// 	log.SetFlags(19)        //(1+2+8)
// 	log.SetPrefix("Mylog:") //前缀
// 	f, err := os.OpenFile("a.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 777)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		log.Fatal("日志文件错误")
// 	}
// 	log.SetOutput(f)
// 	fmt.Printf("log.Flags(): %v\n", log.Flags())
// }
// func log_logger() {
// 	var logger *log.Logger
// 	f, err := os.OpenFile("a.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 777)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		log.Fatal("日志文件错误")
// 	}
// 	logger = log.New(f, "MY log:", 19)
// 	// func log.New(out io.Writer, prefix string, flag int) *log.Logger
// 	logger.Print("get log")
// }
// func main() {
// 	/*
// 		golang内置了log包，实现简单的日志服务。通过调用log包的函数，可以实现简单的日志打印功能。
// 		log包中三个系列的日志打印函数：
// 			1.print：单纯打印日志
// 			2.panic：打印日志，抛出panic异常
// 			3.fatal：打印日志，强制结束程序(os.Exit(1))，defer函数不会执行

// 		const (
// 				Ldate         = 1 << iota     // (1)日期 2009/01/23
// 				Ltime                         // (2)时间 01:23:23
// 				Lmicroseconds                 // (4)微秒时间 01:23:23.123123.  assumes Ltime.
// 				Llongfile                     // (8)完整文件名和行号 /a/b/c/d.go:23
// 				Lshortfile                    // (16)最后的文件名元素和行号 d.go:23. overrides Llongfile
// 				LUTC                          // (32)如果设置了Ldate或Ltime，使用UTC而不是当地时区
// 				Lmsgprefix                    // (64)将“prefix”从行首移动到消息之前
// 				LstdFlags     = Ldate | Ltime // 标准记录器的初始值
// 		)
// 	*/
// 	log_SetFlags()
// 	log.Fatal("get log")
// }
