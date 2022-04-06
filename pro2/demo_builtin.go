// package main

// import (
// 	"fmt"
// )

// func main() {
// 	/*
// 		这个包提供了一些类型声明，还有一些便利函数。
// 	*/
// 	s := []int{}
// 	s = append(s, 1, 2, 3)
// 	fmt.Printf("s: %v\n", s)
// 	s1 := []int{9, 8, 7}
// 	s = append(s, s1...)
// 	fmt.Printf("s: %v\n", s)
// 	s2 := string(append([]byte("hello "), "world"...))
// 	fmt.Printf("s2: %v\n", s2)
// 	// append(slice []Type, elems ...Type) []Type

// 	fmt.Printf("len(s): %v\n", len(s))
// 	fmt.Printf("len(s1): %v\n", len(s1))
// 	fmt.Printf("len(s2): %v\n", len(s2))
// 	// 数组:v中元素的个数。
// 	// 指向数组的指针:*v中的元素数(即使v为nil)。
// 	// 切片或映射:v中元素的数量;如果v为nil, len(v)为0。
// 	// 字符串:v的字节数。
// 	// Channel:通道缓冲区中排队(未读)的元素数;

// 	slice_p := new([]int)
// 	slice := make([]int, 4)
// 	fmt.Printf("slice_p: %T %v\n", slice_p, slice_p)
// 	fmt.Printf("slice: %T %v\n", slice, slice)
// 	// make内置函数分配和初始化类型为slice、map或chan(仅)的对象。new可以分配容易类型的数据。
// 	// 与new一样，第一个参数是类型，而不是值。
// 	// 与new不同，make的返回类型与其参数的类型相同，而不是指向参数的指针。
// 	// new分配的空间被清零，make分配后会进行初始化
// 	// func(Type) *Type
// 	// func(t Type, size ...IntegerType) Type。make slice可以指定两个参数，第一个是长度，第二个是容量，容量>=长度

// 	defer fmt.Println("panic之后还会执行!!!")
// 	panic("GG")
// 	fmt.Println("end...")
// }
