// package main

// import (
// 	"fmt"
// 	"time"
// )

// func show(msg string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(msg)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// }

// func test() {
// 	go show("111")
// 	go show("222")
// 	fmt.Println("end.............")
// }

// func main() {
// 	/*
// 		并发是函数相互独立运行的能力。Goroutines是并发运行的函数。golang提供了goroutines作为并发处理操作的一种方式
// 		创建一个协程非常简单，就是在一个任务函数前面添加一个go关键词
// 		go task()
// 	*/
// 	go show("run...")
// 	show("get...")
// 	fmt.Println("############")
// 	show("run...")
// 	show("get...")
// 	//可以发现前面的方法输出的是"run..."和"get..."混合输出，说明启动了协程

// 	test()
// 	//主协程结束，其他协程都会结束
// 	show("not end") //这条命令让主协程的结束时间拖后了，所以可以多打印一些"111"、"222"
// }
