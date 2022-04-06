// package main

// import (
// 	"fmt"
// )

// type website struct {
// 	Name string
// }

// func main() {
// 	site := website{Name: "kk"}
// 	fmt.Printf("site: %v\n", site)
// 	fmt.Printf("site: %#v\n", site)
// 	fmt.Printf("site: %T\n", site)

// 	fmt.Printf("true: %t\n", true)

// 	/*
// 		二进制
// 		Unicode表示的字符
// 		十进制
// 		八进制
// 		单引号围绕的字符字面值，由go语法转义
// 		十六进制
// 		Unicode格式
// 	*/
// 	fmt.Printf("%b\n", 23)
// 	fmt.Printf("%c\n", 93)
// 	fmt.Printf("%d\n", 23)
// 	fmt.Printf("%o\n", 23)
// 	fmt.Printf("%q\n", 93)
// 	fmt.Printf("%x\n", 23)
// 	fmt.Printf("%U\n", 23)

// 	fmt.Printf("%b\n", 10.2)        //无小数部分的，指数为二的幂的科学计数法
// 	fmt.Printf("%e\n", 10.2)        //科学计数法
// 	fmt.Printf("%E\n", 10.2)        //科学计数法
// 	fmt.Printf("%f\n", 10.00000500) //有小数点而无指数
// 	fmt.Printf("%g\n", 10.2)        //根据情况选择%e%f以产生更紧凑的(无末尾的0)输出

// 	fmt.Printf("%s\n", "hello")
// 	fmt.Printf("%q\n", "hello")
// 	fmt.Printf("%x\n", "Hello")
// 	fmt.Printf("%X\n", "hello")

// 	a := "hello"
// 	p := &a
// 	fmt.Printf("%p\n", p)

// }