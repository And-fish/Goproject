// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func show(i int) {
// 	fmt.Println(i)
// }
// func show_wp(i int) {
// 	defer wp.Done()
// 	fmt.Println(i)
// }

// var wp sync.WaitGroup

// func main() {
// 	// for i := 0; i < 10; i++ {
// 	// 	go show(i)
// 	// }

// 	// fmt.Println("end..........")
// 	// 可以发现go创建的协程还没执行完主程序就结束了，因为主程序结束导致协程也跟着关闭
// 	// 要想控制所有协程结束再进行主程序，需要用到waitgroup

// 	for i := 0; i < 10; i++ {
// 		go show_wp(i)
// 		wp.Add(1)
// 	}
// 	wp.Wait()
// 	fmt.Println("end....")
// 	// 可以发现和上面程序不同的是，for循环仍然是启动协程执行的show_wp()，但是等所有的协程结束了才执行fp(end....)
// 	// wp.add(1)没执行一次就会在后台计数加1，wp.done()执行一次会在后台计数减1
// 	// wp.wait()表示等后台计数为0才会继续执行
// }
