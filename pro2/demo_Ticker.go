// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	/*
// 		timer只执行一次，ticker可以周期执行
// 		ticker1 := time.NewTicker([时间])
// 		遍历返回值是当前时间
// 		ticker1.Reset([时间])	修改周期会新周期
// 		ticker1.stop()		关闭ticker

// 	*/
// 	// count := 1
// 	// ticker1 := time.NewTicker(time.Second * 2)
// 	// for k := range ticker1.C {
// 	// 	if count == 5 {
// 	// 		ticker1.Stop()
// 	// 		break
// 	// 	} else if count == 3 {
// 	// 		ticker1.Reset(time.Second * 5)
// 	// 	}
// 	// 	count++
// 	// 	fmt.Println(k)
// 	// }

// 	ticker1 := time.NewTicker(time.Second)
// 	var channel = make(chan int)
// 	go func() {

// 		for _ = range ticker1.C {
// 			select {
// 			case channel <- 1:
// 			case channel <- 2:
// 			case channel <- 3:
// 			}
// 		}
// 	}()
// 	for v := range channel {
// 		fmt.Printf("v: %v\n", v)
// 	}
// 	// 每隔一秒输入一次
// 	// 这里也正好验证了当select中有多个case可执行时会随机执行其中一个

// }
