// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	a := true
// 	b := false
// 	if a {
// 		fmt.Println(a)
// 	} else {
// 		fmt.Println(b)
// 	}

// 	if x := 20; x > 18 {
// 		fmt.Println("true")
// 	} else {
// 		fmt.Println("false")
// 	}

// 	// var name string
// 	// var age int
// 	// var email string
// 	// fmt.Println("请输入name,age,email")
// 	// fmt.Scan(&name, &age, &email)
// 	// fmt.Printf("name: %v\n", name)
// 	// fmt.Printf("age: %v\n", age)
// 	// fmt.Printf("email: %v\n", email)
// 	// if age%2 == 0 {
// 	// 	fmt.Println("偶数")
// 	// } else {
// 	// 	fmt.Println("奇数")
// 	// }

// 	y := 10
// 	if y > 10 {
// 		fmt.Println("大于10")
// 	} else if y > 5 {
// 		fmt.Println("小于等于10大于5")
// 	} else {
// 		fmt.Println("小于等于5")
// 	}

// 	var day string
// 	fmt.Println("输入第一个字符")
// 	fmt.Scan(&day)
// 	if strings.EqualFold(day, "s") {
// 		fmt.Println("输入第二个字符")
// 		fmt.Scan(&day)
// 		if strings.EqualFold(day, "a") {
// 			fmt.Println("Saturday")
// 		} else if strings.EqualFold(day, "u") {
// 			fmt.Println("Sunday")
// 		} else {
// 			fmt.Println("输入错误")
// 		}
// 	} else if strings.EqualFold(day, "f") {
// 		fmt.Println("Friday")
// 	} else if strings.EqualFold(day, "m") {
// 		fmt.Println("Monday")
// 	} else if strings.EqualFold(day, "t") {
// 		fmt.Println("请输入第二个字符")
// 		fmt.Scan(&day)
// 		if strings.EqualFold(day, "u") {
// 			fmt.Println("Tuesday")
// 		} else if strings.EqualFold(day, "h") {
// 			fmt.Println("Thursday")
// 		} else {
// 			fmt.Println("输入错误")
// 		}
// 	} else if strings.EqualFold(day, "w") {
// 		fmt.Println("Wednesday")
// 	} else {
// 		fmt.Println("输入错误")
// 	}
// }
