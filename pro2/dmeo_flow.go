// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("first")
// 	fmt.Println("secend")

// 	a := true
// 	if a {
// 		fmt.Println("true")
// 	} else {
// 		fmt.Println("false")
// 	}

// 	i := 8
// 	switch {
// 	case i == 10:
// 		fmt.Println("等于10")
// 	case i == 8:
// 		fmt.Println("等于8")
// 		fallthrough
// 		// 如果fallthrough所在的case成立，则下一个无论是否成立都会执行
// 	case i == 5:
// 		fmt.Println("等于5")
// 		fallthrough
// 	default:
// 		fmt.Println("不关心")
// 	}

// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%v\n", i)
// 	}

// 	l := []string{"a", "b", "c", "d"}
// 	for c, v := range l {
// 		fmt.Printf("%v:", c)
// 		fmt.Println(v)

// 	}
// }
