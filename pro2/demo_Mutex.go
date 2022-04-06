// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var i int = 100
// var wg sync.WaitGroup
// var lock sync.Mutex

// func add() {
// 	defer wg.Done()
// 	lock.Lock()
// 	i += 1
// 	lock.Unlock()
// 	// fmt.Printf("i++: %v\n", i)
// }
// func sub() {
// 	defer wg.Done()
// 	lock.Lock()
// 	i -= 1
// 	lock.Unlock()
// 	// fmt.Printf("i--: %v\n", i)
// }
// func main() {
// 	for k := 0; k < 100; k++ {
// 		wg.Add(1)
// 		go add()
// 		wg.Add(1)
// 		go sub()
// 	}
// 	wg.Wait()
// 	// // time.Sleep(10 * time.Second)
// 	fmt.Printf("end i: %v\n", i)

// 	/*
// 		在面对这种两个协程处理一个程序的问题时，要面对两个问题，一个是保证所有处理全部完成，另一个是结果不会出错
// 		1.直接使用sleep等待足够长的时间和waitgroup都可以做到第一个问题，但是不能解决第二个问题
// 		2.所以需要使用锁(mutex)保证同步处理
// 	*/
// }
