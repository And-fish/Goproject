// package main

// import "fmt"

// func main() {
// 	/*
// 		go语言中的函数传参都是值拷贝，当想要改变某个变量时，可以创建一个指向该变量地址的指针变量
// 		传递数据使用指针，而无需拷贝数据
// 		类型指针不能进行偏移和运算
// 		&是取地址，*是根据地址取值

// 		var name *var_type
// 	*/

// 	var intp *int
// 	fmt.Printf("intp: %v\n", intp)
// 	fmt.Printf("intp: %T\n", intp)
// 	var i int = 100
// 	intp = &i
// 	fmt.Printf("intp: %v\n", intp)
// 	fmt.Printf("*intp: %v\n", *intp)

// 	var sp *string
// 	var str string
// 	sp = &str
// 	fmt.Printf("sp: %v\n", sp)
// 	fmt.Printf("sp: %T\n", sp)

// 	a := [...]int{1, 2, 3, 4}
// 	var x int = 1
// 	var pa [len(a) + 1]*int
// 	pa[len(a)] = &x
// 	fmt.Printf("pa: %v\n", pa)
// 	for i := 0; i < len(a); i++ {
// 		pa[i] = &a[i]
// 	}
// 	fmt.Printf("pa: %v\n", pa)

// 	for i := 0; i < len(a); i++ {
// 		fmt.Printf("pa: %v\n", *pa[i])
// 	}
// }
