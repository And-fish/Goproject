// package main

// import "fmt"

// func main() {
// 	/*
// 		go语言中的defer语句会将其后面跟随的语句进行延迟处理。
// 		在defer归属的函数即将返回时，将延迟处理的语句按dafer定义的逆序进行执行，先被defer的语句最后被执行，最后被defer的语句，最先被执行

// 		1.defer用于延迟调用
// 		2.这些调用直到return前才被执行，因此用来做资源清理
// 		3.多个defer语句，按先进后出的方式执行
// 		4.defer语句中的变量，在defer声明时就决定了

// 	*/

// 	fmt.Println("start")
// 	defer fmt.Println("defer1")
// 	defer fmt.Println("defer2")
// 	defer fmt.Println("defer3")
// 	fmt.Println("end")
// 	for i := 0; i < 5; i++ {
// 		defer fmt.Println(i)
// 	}
// 	fmt.Println("all end")

// }
