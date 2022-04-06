// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func show(s string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(s)
// 	}
// }
// func show_goexit() {
// 	for i := 0; i < 20; i++ {
// 		fmt.Printf("i: %v\n", i)
// 		if i == 5 {
// 			runtime.Goexit()
// 		}
// 	}
// }
// func main() {
// 	/*
// 		runtime包里定义了一些协程管理相关的api

// 		runtime.Gosched() 让出CPU时间片，重新安排时间。表示先让其他协程执行
// 			Gosched产生处理器，允许其他goroutines运行。它不会挂起当前的gor例程，所以执行会自动恢复。

// 		runtime.Goexit() 退出当前协程

// 		runtime.GOMAXPROCS(int) 限制核心数，默认是按最大核心数运行，如果设定为1而且不让主程序让出核心就不会优先执行协程的内容
// 	*/

// 	runtime.GOMAXPROCS(1)
// 	go show("golang")
// 	go show_goexit() //可以发现执行到5的时候时候这个协程就退出了
// 	//gosched如果放在这会不一样，个人猜测会至少fp一次"golang"，因为runtime.Gosched()表示在运行这句话的时候让给正在执行的协程，写在这说明只让了一次，能不能执行完show()关键在于show里面的内容执行速度
// 	for i := 0; i < 100; i++ {
// 		runtime.Gosched() //可以发现先执行完了show()再执行了fp(i)，如果把这一行注释掉，show()就会混合在主程序中运行
// 		fmt.Println(i)
// 	}
// 	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
// }
