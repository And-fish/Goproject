// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type name string //类型定义
// type s = string  //类型别名
// func try(s string) bool {
// 	return true
// }
// func main() {
// 	/*
// 		类型定义和类型别名
// 		type name source_type
// 		type name = source_type
// 		类型定义相当于定义了一个全选的类型；但是类型别名只是使用了一个别名来替换之前的类型
// 		定义了类型别名后原来的类型名还能使用
// 		别名只会在代码中存在，在编译完成后不会存在该别名
// 		类型别名可以使用原来类型的方法
// 	*/

// 	var i name
// 	var str string
// 	i = "asdas"
// 	// i = 4567		cannot use 4567 (type untyped int) as type name in assignment
// 	fmt.Printf("i: %T\n", i)
// 	var isstring bool
// 	isstring = reflect.TypeOf(i) == reflect.TypeOf(str)
// 	fmt.Printf("isstring: %v\n", isstring)

// 	var j s = "asda"
// 	isstring = reflect.TypeOf(j) == reflect.TypeOf(str)
// 	fmt.Printf("j: %T\n", j)
// 	fmt.Printf("isstring: %v\n", isstring)

// 	var k string
// 	fmt.Printf("k: %T\n", k)

// 	fmt.Printf("try(k): %v\n", try(k))
// 	fmt.Printf("try(j): %v\n", try(j))
// 	// fmt.Printf("try(i): %v\n", try(i)) 不能运行，因为不是string类型
// }
