// package main

// import "fmt"

// func add_2data(a int, b int) (ans int) {
// 	ans = a + b
// 	return ans
// }
// func subtract(a int, b int) (ans int) {
// 	ans = a * b
// 	return ans
// }

// type calculation func(int, int) int

// func main() {
// 	/*
// 		可以通过一个type关键字去定义一个函数的类型
// 		type name func(p_type) return_type
// 		可以放在主函数里，因为go是编译型语言，所以不能放在其他函数中
// 		???感觉直接把函数赋值到变量上也是可以的，个人感觉这个功能有点鸡肋???
// 	*/

// 	var f calculation
// 	// f_test := subtract
// 	f = subtract
// 	ans := f(1, 2)
// 	// ans_test := f_test(1, 2)
// 	fmt.Printf("ans: %v\n", ans)
// 	// fmt.Printf("ans_test: %v\n", ans_test)
// 	f = add_2data
// 	ans = f(1, 2)
// 	fmt.Printf("ans: %v\n", ans)
// 	fmt.Printf("f: %T\n", f)

// }
