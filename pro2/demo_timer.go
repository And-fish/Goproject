// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	/*
// 		Timer就是定时器的意思，可以实现一些定式操作，内部通过channel来实现的
// 		timer:=time.NewTimer([时间])	||	time.After([时间])
// 	*/
// 	timer1 := time.NewTimer(time.Second * 2)
// 	t1 := time.Now()
// 	fmt.Printf("t1: %v\n", t1)
// 	t2 := <-timer1.C
// 	fmt.Printf("t2: %v\n", t2)
// 	fmt.Println("#################################################################")

// 	timer1 = time.NewTimer(time.Second * 2)
// 	t1 = time.Now()
// 	fmt.Printf("t1: %v\n", t1)
// 	<-timer1.C
// 	t2 = time.Now()
// 	fmt.Printf("t2: %v\n", t2)
// 	fmt.Println("#################################################################")

// 	t1 = time.Now()
// 	fmt.Printf("t1: %v\n", t1)
// 	time.Sleep(time.Second * 2)
// 	t2 = time.Now()
// 	fmt.Printf("t2: %v\n", t2)
// 	fmt.Println("#################################################################")

// 	t1 = time.Now()
// 	fmt.Printf("t1: %v\n", t1)
// 	<-time.After(time.Second * 2)
// 	t2 = time.Now()
// 	fmt.Printf("t2: %v\n", t2)
// 	fmt.Println("#################################################################")

// 	timer1 = time.NewTimer(time.Hour)
// 	// timer1 := time.NewTimer(time.Second)
// 	go func() {
// 		<-timer1.C
// 	}()
// 	// time.Sleep(time.Second * 3)
// 	stop := timer1.Stop()
// 	// 如果调用stop则返回true，如果计时器已经停止，则返回false
// 	// 声明即调用
// 	if stop {
// 		fmt.Println(stop)
// 	} else {
// 		fmt.Println(stop)
// 	}
// }
