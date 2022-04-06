// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// 	"time"
// )

// var i int = 100
// var j int64 = 100

// func add() {
// 	i++
// }
// func sub() {
// 	i--
// }
// func a_add() {
// 	atomic.AddInt64(&j, 1)
// }
// func a_sub() {
// 	atomic.AddInt64(&j, -1)
// }

// func main() {
// 	/*
// 		atomic提供的原子操作能过确保任意时刻只有一个goroutine对变量进行操作，善用atomic能过避免程序出现大量的锁行为
// 		atomic常见操作：
// 			1.增加
// 			2.载入
// 			3.比较并交换cas
// 			4.交换
// 			5.存储
// 	*/

// 	for i := 0; i < 100; i++ {
// 		go add()
// 		go sub()
// 	}
// 	time.Sleep(time.Second * 3)
// 	fmt.Printf("i: %v\n", i)

// 	for i := 0; i < 100; i++ {
// 		go a_add()
// 		go a_sub()
// 	}
// 	time.Sleep(time.Second * 3)
// 	fmt.Printf("i: %v\n", j)
// 	/*
// 		增减:
// 			atomic.AddInt32()
// 			atomic.AddInt64()
// 			atomic.AddUint32()
// 			atomic.AddUint64()
// 			atomic.AddUintptr()
// 	*/

// 	var k int32 = 100
// 	atomic.LoadInt32(&k)
// 	atomic.StoreInt32(&k, 200)
// 	fmt.Printf("k: %v\n", k)
// 	/*
// 		读写：
// 			载入能够保证院子的读变量的值，当读取的时候，如何其他的cpu操作都无法对该变量进行读写。这也是atomic实现的核心
// 			store确保了协变量的原子性，避免其他操作读到了修改变量过程中的脏数据
// 	*/

// 	var n int32 = 100
// 	isok := atomic.CompareAndSwapInt32(&n, 100, 200)
// 	// func atomic.CompareAndSwapInt32(addr *int32, old int32, new int32) (swapped bool)
// 	fmt.Printf("isok: %v\n", isok)
// 	fmt.Printf("n: %v\n", n)
// 	isok = atomic.CompareAndSwapInt32(&n, 100, 300)
// 	// func atomic.CompareAndSwapInt32(addr *int32, old int32, new int32) (swapped bool)
// 	fmt.Printf("isok: %v\n", isok)
// 	fmt.Printf("n: %v\n", n)
// 	/*
// 		对比交换：
// 				如果n==old int，就把new int赋值给n
// 				否则就不交换，这是atomic实现的核心
// 	*/

// 	var m int32 = 100
// 	atomic.AddInt32(&m, 200)
// 	fmt.Printf("m: %v\n", m)
// 	/*
// 		交换:
// 			相比于CAS，这个操作更加暴力，不管变量的旧值是否被改变，直接赋予新值然后返回被替换的值
// 	*/
// }
