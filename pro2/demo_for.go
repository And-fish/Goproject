// package main

// import (
// 	"fmt"
// )

// func main() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("i: %v ", i)
// 	}

// 	fmt.Println("")
// 	a := 0
// 	for ; a < 10; a++ {
// 		fmt.Printf("a: %v ", a)
// 	}

// 	fmt.Println("")
// 	b := 0
// 	for b < 10 {
// 		fmt.Printf("b: %v ", b)
// 		b++
// 	}

// 	fmt.Println("")
// 	c := 0
// 	for {
// 		if c < 10 {
// 			fmt.Printf("c: %v ", c)
// 			c++
// 			continue
// 		}
// 		break
// 	}

// 	fmt.Println("")
// 	var x = [...]int{1, 2, 3, 4, 5}
// 	//数组
// 	for i, v := range x {
// 		fmt.Printf("i: %v ", i)
// 		fmt.Printf("v: %v\n", v)
// 	}

// 	fmt.Println(" ")
// 	var y = []int{1, 2, 3}
// 	//切片
// 	for _, v := range y {
// 		fmt.Printf("v: %v\n", v)
// 	}

// 	fmt.Println(" ")
// 	var m = make(map[string]string, 0)
// 	m["name"] = "tom"
// 	m["age"] = "20"
// 	m["email"] = "123@qq.com"
// 	for key, va := range m {
// 		fmt.Printf("key: %v ", key)
// 		fmt.Printf("va: %v\n", va)
// 	}

// 	s := "hello world"
// 	for _, v := range s {
// 		fmt.Printf("v: %q ", v)
// 	}
// }
