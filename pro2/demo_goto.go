// package main

// import "fmt"

// func main() {
// 	i := 1
// 	if i > 2 {
// 		fmt.Println("true")
// 	} else {
// 		goto end
// 	}
// end:
// 	fmt.Println("end")

// 	for i := 0; i < 5; i++ {
// 		for j := 0; j < 5; j++ {
// 			if i >= 2 && j >= 2 {
// 				goto beybey
// 			}
// 			fmt.Println(i, j)
// 		}
// 	}
// beybey:
// 	fmt.Println("beybey")

// 	j := 1
// next:
// 	j++
// 	if j > 5 {
// 		fmt.Printf("j: %v\n", j)
// 		fmt.Println("true")
// 	} else {
// 		goto next
// 	}

// 	//建议把goto到的label写到goto语句的后面，否则很容易混乱逻辑
// }
