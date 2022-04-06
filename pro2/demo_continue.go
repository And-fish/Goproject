// package main

// import "fmt"

// func main() {
// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			continue
// 		}
// 		fmt.Printf("i: %v ", i)
// 	}
// 	fmt.Println("")

// 	for i := 0; i < 5; i++ {
// 	label:
// 		for j := 0; j < 5; j++ {
// 			if i == j && i == 2 {
// 				continue label
// 			}
// 			fmt.Printf("%v,%v ", i, j)
// 		}
// 	}

// 	/*
// 		label的使用:
// 			continue:跳转到label所在循环的本次循环
// 			break:结束label所在的循环
// 		不加label:
// 			continue:跳过本次循环
// 			break:结束本层循环
// 	*/
// }
