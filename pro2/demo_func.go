// package main

// import (
// 	"fmt"
// )

// func main() {
// 	/*
// 		go里面有三种函数：普通函数、匿名函数、方法（定义在struct上的函数）
// 		go语言中不允许重载(overload)，也就是说不允许函数同名
// 		go语言中的函数不能嵌套函数，但是可以嵌套匿名函数
// 		函数是一个值，可以将函数赋值给变量，使得这个变量也成为函数
// 		函数可以作为参数传递给另一个函数
// 		函数的返回值可以是一个函数
// 		函数调用的时候如果有参数传递给函数，则先拷贝参数的副本，再将副本传递给函数
// 		函数参数可以没有名称

// 		// parameter  参数
// 		func name(parameter p_type)(return_value return_type){
// 			函数体
// 		}
// 		return的值可以不是return声明里的value，但是要在函数体中重新声明而且type要一致
// 		可以函数体中如果没有return，可以不写return声明

// 		如果传递的参数是map、slice、interface、channel这些数据类型是可能会改变的。原因和之前切片中写的一样，这些数据类型本质是是指针，传递的时候指针拷贝了仍然是指向着底层的数据
// 		变长参数，在声明参数的时候在参数名字后面加上“...”表示不限制数量。一个函数只能拥有一个，而且只能放在最后，传递的是一个切片

// 	*/
// 	sum_ans := get_sum(1, 2)
// 	fmt.Printf("a: %v\n", sum_ans)

// 	func1()
// 	func2("func2")
// 	fmt.Printf("func3(\"func3\"): %v\n", func3("func3"))
// 	fmt.Printf("func4(\"func4\"): %v\n", func4("func4"))

// 	f5_1, f5_2 := func5(5, 10)
// 	fmt.Println(f5_1, f5_2)

// 	slice1 := []int{1, 2, 3, 4}
// 	var slice1_copy = make([]int, len(slice1))
// 	copy(slice1_copy, slice1)
// 	func6(slice1_copy)
// 	issame := is_sameslice(slice1, slice1_copy)
// 	fmt.Printf("issame: %v\n", issame)

// 	func7("sa", 1, 2, 3, 4)
// }

// func get_sum(a int, b int) (ans int) {

// 	ans = a + b
// 	return ans
// }

// func is_sameslice(slice1 []int, slice1_b []int) (ans bool) {
// label:
// 	for {
// 		for i := 0; i < len(slice1); i++ {
// 			if slice1[i] != slice1_b[i] {
// 				ans = false
// 				break label
// 			}
// 		}
// 		ans = true
// 		break label
// 	}
// 	return ans
// }

// func func1() {
// 	fmt.Println("没有参数也没有返回值")
// }
// func func2(a string) {
// 	fmt.Println("只有参数没有返回值")
// }
// func func3(a string) (ret string) {
// 	ret = "又有参数又有返回值"
// 	return ret
// }
// func func4(a string) string {
// 	ret := "又有参数又有返回值但是没有直接声明变量名称"
// 	return ret
// }
// func func5(a int, b int) (k int, t int) {
// 	k = a
// 	t = b
// 	return
// }
// func func6(a []int) {
// 	a[0] = 100
// }
// func func7(b string, a ...int) {
// 	fmt.Println(b)
// 	fmt.Printf("%T	", a)
// }
