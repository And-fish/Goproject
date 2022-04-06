// package main

// import "fmt"

// type Human struct {
// 	name     string
// 	age      int
// 	birthday string
// }

// func (a Human) eat() {
// 	fmt.Println(a.name, ",eat...")
// }
// func (a Human) sleep() {
// 	fmt.Println(a.name, ",sleep...")
// }
// func (a Human) work() {
// 	fmt.Println(a.name, ",work...")
// }

// type family struct {
// 	counting int
// 	father   Human
// 	mother   Human
// 	kid      Human
// 	int
// }

// func main() {
// 	/*
// 		利用结构体和函数绑定来实现oop
// 		利用结构体嵌套来实现类的继承
// 	*/
// 	xxh := Human{
// 		name:     "xxh",
// 		age:      19,
// 		birthday: "2002/08/20",
// 	}
// 	xxh.eat()
// 	xxh.sleep()
// 	xxh.work()
// 	fmt.Println("#######################")
// 	mom := Human{name: "lxf"}
// 	dad := Human{name: "xh"}
// 	var F1408 family
// 	F1408 = family{father: dad, mother: mom, kid: xxh, counting: 3}
// 	fmt.Printf("F1408: %v\n", F1408)

// 	F1408.father.work()

// }
