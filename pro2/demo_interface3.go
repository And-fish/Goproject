// package main

// import "fmt"

// type flyfish interface {
// 	fly
// 	swim
// }
// type fly interface {
// 	fly()
// }
// type swim interface {
// 	swim()
// }

// type fish struct {
// 	id int
// }

// func (fish fish) swim() {
// 	fmt.Println("swimming...")
// }
// func (fish fish) fly() {
// 	fmt.Println("flying...")
// }
// func main() {
// 	/*
// 		inferface的嵌套，例如飞鱼，又可以飞又可以游泳
// 	*/
// 	fish1 := fish{1}
// 	fish1.fly()
// 	fish1.swim()

// 	var fish2 fly
// 	fish2 = fish{2}
// 	fish2.fly()
// 	// fish2.swim() 报错，因为fly中没有swim的方法

// 	var fish3 swim
// 	fish3 = fish{3}
// 	// fish3.fly() 报错，因为swim中没有fly的方法
// 	fish3.swim()

// 	var fish4 flyfish
// 	fish4 = fish{4}
// 	fish4.fly()
// 	fish4.swim() //成功，因为flyfish中包括了fly和swim接口，而且fly和swim接口中有对应的方法
// }
