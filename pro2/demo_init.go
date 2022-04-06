// package main

// import "fmt"

// func init() {
// 	fmt.Println("init...")
// }

// func init() {
// 	fmt.Println("init2...")
// }

// var test_var int = getvar()

// func getvar() int {
// 	fmt.Println("var...")
// 	return 100
// }
// func main() {
// 	/*
// 		init函数先于main函数执行，用于实现包级别的一些初始化
// 		初始化顺序：变量初始化-->init()-->main()
// 		一个包可以有多个init函数
// 	*/

// 	fmt.Println("main...")
// 	//可以发现先输出了var...再输出了init...\init...再输出main...

// }
