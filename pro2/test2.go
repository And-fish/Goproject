// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func run(s string){
// 	fmt.Println("runn")
// }
// func main() {
// 	var m = map[string]interface{}{"name": "hk", "age": 1999}
// 	for k, v := range m {
// 		iss := reflect.TypeOf(v) == reflect.TypeOf(" ")
// 		fmt.Printf("iss: %v\n", iss)
// 		switch v.(type) {
// 		case string:
// 			fmt.Printf("k: %v\n", k)
// 			run(v)
// 		case int:
// 			fmt.Printf("k1: %v\n", k)
// 		default:
// 			fmt.Printf("k2: %v\n", k)
// 		}
// 	}
// }
package main

import (
	"fmt"
)

func main() {
	fmt.Sprintf("%v")
}
