// package main

// import (
// 	"fmt"
// )

// type Human struct {
// 	name  string
// 	doing Human_doing
// }
// type Human_doing func()

// func (a Human) eat() {
// 	fmt.Printf("%v,eating...\n", a.name)
// }
// func (a Human) sleep() {
// 	fmt.Printf("%v.sleeping...\n", a.name)
// }

// func (a Human) changename_v() {
// 	a.name = "v_changed"
// 	fmt.Printf("a.name: %v\n", a.name)
// }
// func (p *Human) changename_p() {
// 	p.name = "p_changed"
// 	fmt.Printf("p.name: %v\n", p.name)
// }

// func kiss() {
// 	fmt.Println("kiss..")
// }
// func main() {
// 	/*
// 		我们可以使用结构体来模拟面向对象和类对象。我们都知道面向对象里面有类方法的该娘，我们也可以声明一些方法，属于某个结构体
// 		go语言中的方法是一种特殊的函数，定义与struct之上，被称为struct的接受者(receive)
// 		方法就是有接受者的函数
// 		type struct_name struct{}
// 		func (recv) method_name(参数) return_type{}
// 		recv可以是结构体，也可以是结构体指针
// 		可以看出一个方法和一个函数非常相似，多了一个接受类型
// 		//就像是class下面的def，可以直接用"."来调用
// 	*/
// 	human1 := Human{name: "human1"}
// 	human1.eat()
// 	human1.sleep()

// 	var f Human_doing
// 	f = human1.eat
// 	f()
// 	fmt.Printf("human1: %p\n", &human1)
// 	fmt.Printf("f: %v\n", &f)

// 	human2 := Human{name: "human2"}

// 	human1.doing = human2.eat
// 	human1.doing()
// 	human1.doing = kiss
// 	human1.doing()
// 	//可以定义结构体的时候添加函数

// 	/*
// 		1.接受者的类型不一定是struct，还可以是类型别名、slice、map、func等等
// 		2.struct结合它的方法，可以和方法分开，并非要属于一个文件，但是必须属于一个包
// 		3.方法有两种接受类型(Type)和(*Type)
// 		4.方法就是函数，因为go语言不能重载，所以不能重名
// 		5.如果接受者是一个指针类型，会自动解除引用
// 		6.方法和type是分开的，意味着实例的行为和数据存储是分开的，但是他们通过receive建立起关联关系
// 	*/

// 	/*
// 		方法接受者类型：
// 			结构体实例有，值类型和指针类型；那么方法的接受者也有值类型和指针类型
// 		根据之前的经验，不妨猜测，当使用值类型方法的时候不会导致外部的内容变化，但是使用指针类型会改变内容
// 	*/
// 	human3 := Human{name: "human3"}
// 	fmt.Printf("human3.name: %v\n", human3.name)
// 	human3.changename_v()
// 	fmt.Printf("human3.name: %v\n", human3.name)
// 	human3.changename_p()
// 	fmt.Printf("human3.name: %v\n", human3.name)

// 	fmt.Println("#################")
// 	p := Human{name: "human4"}
// 	p1 := &p
// 	fmt.Printf("p: %v\n", p)
// 	fmt.Printf("p1: %v\n", p1)

// 	p.changename_v()
// 	fmt.Printf("p: %v\n", p)
// 	p1.changename_v()
// 	fmt.Printf("p1: %v\n", p1)

// 	p.changename_p()
// 	fmt.Printf("p: %v\n", p)
// 	fmt.Printf("p1: %v\n", p1)
// 	p.name = "human4"
// 	p1.changename_p()
// 	fmt.Printf("p: %v\n", p)
// 	fmt.Printf("p1: %v\n", p1)
// 	// 猜想正确，这次猜想进一步证实了底层参数传递的规律
// 	// 而且可以发现，接受者的类型是指针类型并不会让外部内容发生变化
// 	// 只要使用的是指针类型的方法就会修改外部的内容
// }
