// package main

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// func errors_test(s string) error {
// 	if s == "" {
// 		return errors.New("字符串不能为空")
// 	}
// 	return nil
// }
// func main() {
// 	/*
// 		errors包实现了操作错误的函数。语言使用error类型来返回函数执行过程汇总遇到的错误。
// 		type error interface {
// 			Error() string
// 		}
// 		errors包实现了一个最简单的error类型，只包含一个字符串。使用New函数，用于生成一个最简单的error对象
// 	*/

// 	err := errors_test("")
// 	fmt.Printf("err: %T\n", err)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	err1 := oops()
// 	fmt.Printf("err1: %T\n", err1)
// 	if err1 != nil {
// 		fmt.Printf("err1: %v\n", err1)
// 	} else {
// 		fmt.Println(err1)
// 	}

// }

// type MyError struct {
// 	when time.Time
// 	what string
// }

// func (e MyError) Error() string {
// 	return fmt.Sprintf("%v %v", e.when, e.what)
// }
// func oops() error {
// 	return &MyError{
// 		time.Now(),
// 		"get err",
// 	}
// }
