// package main

// import (
// 	"fmt"
// )

// var channelint = make(chan int)
// var channelstr = make(chan string)

// func main() {
// 	/*
// 		1.select是go的一个控制结构，类似于switch语句，用于处理异步IO操作。
// 		  select会简体case语句中channel的读写操作，当case中channel读写操作为非阻塞状态(即能读写)时，将会出发相应动作
// 				select中的case语句必须是一个channel操作
// 				select中的default子句总是可运行的
// 		2.如果有多个case都可以运行，select会随机公平地选出一个执行，其他的不会执行
// 		3.如果没有可运行case语句，且有default语句，那么就会执行default的动作
// 		4.如果没有可运行case语句，且没有default语句，select将阻塞，直到某个case通信可以运行

// 	*/
// 	go func() {
// 		channelint <- 100
// 		fmt.Println("100已经传输")
// 		channelint <- 200
// 		fmt.Println("200已经传输") //再添加了这两句之后有无缓冲的区别变得更加明显
// 	}()

// 	// time.Sleep(time.Second * 3)
// 	fmt.Printf("len(channelint): %v\n", len(channelint))
// 	// 这里说一下有缓冲通道和无缓冲通道的区别，通过上面这条可以发现，
// 	// 当channelint是无缓冲通道的时候，因为我还没接受，所以此时通道内是没有东西的
// 	// 说的可能过于抽象，但是当我把channelint改成有缓冲通道，就可以发现区别。
// 	// 修改之后，我还没有接收指令，通道内就已经被传入了数据

// 	/*
// 		无缓冲:		因为没有执行到接收数据，所以在time.sleep的时间里通道内没有传入数据
// 			len(channelint): 0
// 			data: 100
// 			100已经传输
// 			200已经传输
// 			data: 200

// 		有缓冲:		可以发现，在time.sleep的时间里，通道内已经被传入了数据
// 			100已经传输
// 			200已经传输
// 			len(channelint): 2
// 			data: 100
// 			data: 200
// 	*/
// 	data := <-channelint
// 	fmt.Printf("data: %v\n", data)
// 	data = <-channelint
// 	fmt.Printf("data: %v\n", data)

// 	go func() {
// 		// defer close(channelint)
// 		// defer close(channelstr)
// 		channelint <- 100
// 		channelstr <- "hello"
// 	}()
// label:
// 	for {
// 		select {
// 		case data = <-channelint:
// 			fmt.Printf("data: %v\n", data)
// 		case value := <-channelstr:
// 			fmt.Printf("value: %v\n", value)
// 		default:
// 			fmt.Println("gg")
// 			break label
// 		}
// 	}

// }
