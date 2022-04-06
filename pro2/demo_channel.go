// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// var intchannel = make(chan int)

// func send() {
// 	defer close(intchannel)
// 	rand.Seed(time.Now().UnixNano())
// 	value := rand.Intn(10)
// 	intchannel <- value
// }

// func main() {
// 	/*
// 		go提供了一种称为通道的机制，用于goroutine之间共享数据。
// 		当goroutine执行并发活动需要共享资源或数据时，通道提供了一种机制保证同步交换
// 		声明通道时需要指定数据类型
// 		有两种类型的通道：
// 			1.无缓冲通道：用于执行同步通信，保证发生和接受发生的瞬间执行两个程序的交换
// 						Ubbuffered:=make(chan int)		//整型无缓冲通道
// 			2.有缓冲通道：用于执行异步通信。
// 						buffered:=make(chan int,10)		//整型有缓冲通道
// 	*/

// 	go send() //!!!必须是两个goroutine之间的交流才用通道，所以必须用go关键词启动一个协程!!!
// 	fmt.Println("wait")
// 	data := <-intchannel
// 	fmt.Printf("data: %v\n", data)
// 	fmt.Println("end...")

// 	intchannel = make(chan int)
// 	go func() {
// 		for i := 0; i < 2; i++ {
// 			intchannel <- i
// 		}
// 		close(intchannel) // 如果不关，则不能读超过写的次数
// 	}()
// 	time.Sleep(time.Second * 2)
// 	for i := 0; i < 5; i++ {
// 		data = <-intchannel
// 		fmt.Printf("data: %v\n", data)
// 	}

// 	intchannel = make(chan int)
// 	go func() {
// 		for i := 0; i < 2; i++ {
// 			intchannel <- i
// 		}
// 		close(intchannel)
// 	}()
// 	time.Sleep(time.Second * 2)
// 	for value := range intchannel {
// 		fmt.Printf("value: %v\n", value)
// 	}

// 	intchannel = make(chan int)
// 	go func() {
// 		for i := 0; i < 2; i++ {
// 			intchannel <- i
// 		}
// 		close(intchannel)
// 	}()
// 	time.Sleep(time.Second * 2)
// 	for {
// 		data, ok := <-intchannel
// 		if ok {
// 			fmt.Printf("data: %v\n", data)
// 		} else {
// 			break
// 		}
// 	}
// }
