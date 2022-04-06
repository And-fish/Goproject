// package main

// import (
// 	"fmt"
// )

// func sayhello(name string) {
// 	fmt.Printf("hello,%v\n", name)
// }

// func func1(name string, other_func func(string)) {
// 	other_func(name)
// }
// func run(name string) {
// 	func1(name, sayhello)
// }

// type do func(int, int) int

// func add_2data(a int, b int) (ans int) {
// 	ans = a + b
// 	return ans
// }
// func subtract(a int, b int) (ans int) {
// 	ans = a * b
// 	return ans
// }
// func err_do(a int, b int) int {
// 	return 0
// }
// func how2do(a string) func(int, int) int { //如果写出func how2do(a string) get就很容易理解了
// 	switch {
// 	case a == "+":
// 		return add_2data
// 	case a == "*":
// 		return subtract
// 	default:
// 		return err_do
// 	}
// }

// func always_add() func(int) int {
// 	var x int
// 	return func(i int) int {
// 		x += i
// 		return x
// 	}
// }

// func main() {
// 	/*
// 		函数可以作为函数的参数，也可以作为另一个函数的返回值
// 	*/
// 	sayhello("xxh")

// 	run("xxh")

// 	do := how2do("+")
// 	fmt.Printf("do(1, 2): %v\n", do(1, 2))

// 	/*
// 		匿名函数:因为go语言中不能嵌套，但是在函数内部可以定义匿名函数，实现功能调用
// 		func (p p_type)(return_value return_type){
// 			函数体
// 		}
// 		在匿名函数的后面加上(参数)会自己执行
// 	*/

// 	getmax := func(a int, b int) int {
// 		if a > b {
// 			return a
// 		} else {
// 			return b
// 		}
// 	}

// 	themax := getmax(1, 2)
// 	fmt.Printf("themax: %v\n", themax)

// 	d1 := func(a int, b int) int {
// 		if a > b {
// 			return a
// 		} else {
// 			return b
// 		}
// 	}(5, 4)
// 	fmt.Printf("d1: %v\n", d1)

// 	/*
// 		闭包函数可以理解为定义在一个函数内部的函数
// 	*/
// 	f := always_add()
// 	// f==func(i int) int {
// 	// 	x += i
// 	// 	return x
// 	// }
// 	// 同时定义了一个x在内存中，而且因为这时候f()没有重新定义新的x，所以后续调用f()只会修改x
// 	r := f(10)
// 	fmt.Printf("r: %v\n", r)
// 	r = f(20)
// 	fmt.Printf("r: %v\n", r) //r==30,
// 	f2 := always_add()
// 	// f2==func(i int) int {
// 	// 	x += i
// 	// 	return x
// 	// }
// 	// 同时定义了一个x2在内存中，而且因为这时候f2()没有重新定义新的x2，所以后续调用f()只会修改x2
// 	r = f2(50)
// 	fmt.Printf("r: %v\n", r)
// 	r = f(100)
// 	fmt.Printf("r: %v\n", r)
// 	f = always_add()
// 	// f==func(i int) int {
// 	// 	x += i
// 	// 	return x
// 	// }
// 	// 同时重新定义了一个x3在内存中，因为f()已经不是原来的x3，所以调用f()会从x刚刚定义开始
// 	r = f(500)
// 	fmt.Printf("r: %v\n", r)

// 	fx, fy := todo(10)
// 	fmt.Println(fx(2), fy(1))
// 	fmt.Println(fx(8), fy(3))
// 	// 这里的todo函数传的参数和上一个闭包函数的声明x起到了一样的作用，
// 	// 因为todo返回的两个匿名函数都不会重新定义一个数据，只会在原来a的基础上修改，除非重新定义fx和fy

// 	fx = func(i int) int {
// 		_ = i
// 		return 0
// 	}
// 	fmt.Println(fx(8), fy(10))
// 	// 可以发现我重新定义了fx，后续调用fx自然会执行新定义的fx，但是fy仍然对着a执行函数。这意味着a1并不会随着重新声明而释放掉内存

// 	fx, fy = todo(10)
// 	fmt.Println(fx(2), fy(1))
// 	fmt.Println(fx(8), fy(3))
// 	//个人猜测这里虽然重新声明了fx和fy，a1仍然在内存中等待原来的fx和fy使用，但是早已物是人非，fx、fy再也不可能变成原来模样，他们永远离开了她

// 	fx2, fy2 := todo(100)
// 	fmt.Println(fx2(2), fy2(1))
// 	fmt.Println(fx2(8), fy2(3))

// 	/*
// 		函数内部调用函数自身的函数称为递归函数。
// 		使用递归函数最重要的三点：
// 			1.递归就是自己调用自己
// 			2.必须先定义函数的退出条件，没有退出条件，递归将成为死循环
// 			3.go语言递归函数很可能会产生一大堆的goroutine，也很可能会出现栈空间内存溢出问题
// 	*/
// 	ans := recursive_f1(1)
// 	fmt.Printf("ans: %v\n", ans)

// 	ans = recursive_f2(5)
// 	fmt.Printf("ans: %v\n", ans)
// }

// func todo(a int) (func(int) int, func(int) int) {
// 	add := func(i int) int {
// 		a += i
// 		return a
// 	}

// 	sub := func(i int) int {
// 		a -= i
// 		return a
// 	}
// 	return add, sub
// }

// func recursive_f1(n int) int {
// 	// 查看每次加1，从1加到10的次数
// 	if n == 10 {
// 		return 0
// 	}
// 	n = n + 1
// 	ans := recursive_f1(n) + 1
// 	return ans
// }

// func recursive_f2(n int) int {
// 	if n == 1 {
// 		return 1
// 	}

// 	ans := recursive_f2(n-1) * n
// 	return ans
// }
