// package main

// import (
// 	"fmt"
// )

// type Human struct {
// 	name string
// 	age  int
// }

// func NewHuman(name string, age int) (*Human, error) {
// 	if name == "" {
// 		return nil, fmt.Errorf("name不能为空")
// 	} else if age < 0 {
// 		return nil, fmt.Errorf("请输入正确的年龄")
// 	} else {
// 		p := new(Human)
// 		p.name = name
// 		p.age = age
// 		return p, nil
// 	}
// 	/*
// 		nil表示指针、通道、函数、接口、映射、切片
// 		不能表示结构体，所以返回的时候选择返回*Human而不是Human
// 	*/

// }
// func main() {
// 	h1, e1 := NewHuman("xxh", 20)
// 	fmt.Printf("h1: %v,e1:%v\n", h1, e1)
// 	h1, e1 = NewHuman("", 20)
// 	fmt.Printf("h1: %v,e1:%v\n", h1, e1)
// 	h1, e1 = NewHuman("xxh", -5)
// 	fmt.Printf("h1: %v,e1:%v\n", h1, e1)
// }
