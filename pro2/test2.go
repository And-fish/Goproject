// // package main

// // import (
// // 	"fmt"
// // 	"reflect"
// // )

// // func run(s string){
// // 	fmt.Println("runn")
// // }
// // func main() {
// // 	var m = map[string]interface{}{"name": "hk", "age": 1999}
// // 	for k, v := range m {
// // 		iss := reflect.TypeOf(v) == reflect.TypeOf(" ")
// // 		fmt.Printf("iss: %v\n", iss)
// // 		switch v.(type) {
// // 		case string:
// // 			fmt.Printf("k: %v\n", k)
// // 			run(v)
// // 		case int:
// // 			fmt.Printf("k1: %v\n", k)
// // 		default:
// // 			fmt.Printf("k2: %v\n", k)
// // 		}
// // 	}
// // }
// // package main

// // import "fmt"

// // func main() {
// // 	var array = [5]int{3, 1, 4, 5, 2}
// // 	for i, v := range array {
// // 		for _, u := range array[i+1:] {
// // 			if v > u {
// // 				fmt.Printf("[]int{v, u}: %v\n", []int{v, u})
// // 			}
// // 		}
// // 	}
// // }
// package main

// import "fmt"

// func do(array []int) {
// 	for i, v := range array {
// 		fmt.Printf("i: %v\n", i)
// 		fmt.Printf("v: %v\n", v)
// 	}
// }
