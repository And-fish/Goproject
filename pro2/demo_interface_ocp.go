// package main

// import "fmt"

// type pet interface {
// 	eat()
// 	sleep()
// }

// type Human struct {
// 	name string
// }
// type dog struct {
// 	name string
// }
// type cat struct {
// 	name string
// }

// func (human Human) keep(a pet) {
// 	a.eat()
// 	a.sleep()
// }
// func (a dog) eat() {
// 	fmt.Printf("%v,", a.name)
// 	fmt.Println("eat...")
// }
// func (a dog) sleep() {
// 	fmt.Printf("%v,", a.name)
// 	fmt.Println("sleep")
// }
// func (a cat) eat() {
// 	fmt.Printf("%v,", a.name)
// 	fmt.Println("eat...")
// }
// func (a cat) sleep() {
// 	fmt.Printf("%v,", a.name)
// 	fmt.Println("sleep")
// }
// func main() {
// 	/*
// 		opc设计原则:Open-Closed Principle“开-闭”原则
// 		对扩展是开放的，对修改是关闭的
// 	*/
// 	dog1 := dog{"lupi"}
// 	cat1 := cat{"kiti"}
// 	human1 := Human{"tom"}
// 	human1.keep(dog1)
// 	cat1.sleep()

// }
