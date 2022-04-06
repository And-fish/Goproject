// package main

// import (
// 	"fmt"
// )

// type Human struct {
// 	id    int
// 	name  string
// 	age   int
// 	email string
// }

// func main() {
// 	/*
// 		go语言没有面向对象的概念，但是可以用结构体来实现面向对象的一节特性，例如继承、组合等特性
// 		结构体的定义：
// 			type name struct{
// 				成员定义
// 			}
// 	*/

// 	type animal struct {
// 		length, weight int
// 		name           string
// 	}

// 	var tom Human
// 	fmt.Printf("tom: %v\n", tom)
// 	fmt.Printf("tom: %T\n", tom)

// 	kit := Human{}
// 	fmt.Printf("kit: %v\n", kit)
// 	// 没有声明结构体中属性的话会自动补充（int:0,string:" "）

// 	var human1 Human
// 	human1.id = 1
// 	human1.name = "human1"
// 	human1.age = 18
// 	human1.email = "1@gmail.com"
// 	fmt.Printf("human1: %v\n", human1)
// 	var human2 = Human{id: 2, name: "human2", age: 19, email: "2@gmail.com"}
// 	fmt.Printf("human2: %v\n", human2)
// 	human3 := Human{id: 3, name: "human3", age: 16, email: "3@gmail.com"}
// 	fmt.Printf("human3: %v\n", human3)
// 	human4 := Human{4, "human4", 20, "4@gmail.com"}
// 	fmt.Printf("human4: %v\n", human4)
// 	human5 := Human{
// 		5,
// 		"human5",
// 		30,
// 		"5@gmail.com", //注意这里有一个逗号
// 	}
// 	fmt.Printf("human5: %v\n", human5)

// 	person := animal{173, 50, "person"}
// 	person = animal{weight: 50, length: 173, name: "person"}
// 	fmt.Printf("person: %v\n", person)

// 	/*
// 		匿名结构体：
// 		var name struct{
// 			成员定义
// 		}
// 	*/
// 	var dog struct {
// 		id   int
// 		name string
// 	}
// 	dog.id = 100
// 	dog.name = "kk"
// 	fmt.Printf("dog: %v\n", dog)

// 	/*
// 		结构体指针
// 	*/
// 	p_human1 := &human1
// 	fmt.Printf("p_human1: %v\n", p_human1)
// 	fmt.Printf("p_human1: %p\n", p_human1)
// 	fmt.Printf("p_human1: %v\n", *p_human1)

// 	p_human2 := new(Human)
// 	fmt.Printf("p_human2: %v\n", p_human2)
// 	p_human2 = &human2
// 	fmt.Printf("p_human2: %v\n", p_human2)

// 	var p_human3 *Human
// 	p_human3 = &human3
// 	fmt.Printf("p_human3: %v\n", p_human3)

// 	p_human4 := new(Human)
// 	fmt.Printf("p_human4: %v\n", p_human4)
// 	fmt.Printf("p_human4: %p\n", p_human4)
// 	p_human4.id = 10
// 	(*p_human4).age = 6 //可以省略*
// 	fmt.Printf("p_human4: %v\n", p_human4)
// 	fmt.Printf("p_human4: %p\n", p_human4)

// 	/*
// 		结构体可以作为函数的参数:
// 			1.直接传递个结构体，是一个copy，在函数内部不会改变外面的结构
// 			2.如果传递的是结构体的指针，会改变外部结构体的内容
// 			//个人认为这就是为什么把数组（切片）和int(string、float64)传进函数
// 			//一个会修改外部参数一个不会的原因
// 			//切片在传递的时候本质上是直接传递的地址而不是切片本身
// 	*/

// 	human7 := Human{7, "human7", 50, "7@gmail.com"}
// 	fmt.Printf("human7: %v\n", human7)
// 	ischange_m(human7)
// 	fmt.Printf("human7: %v\n", human7)
// 	//可以发现在函数内结构体内容发生了改变，但是改变值体现在函数内部
// 	p_human7 := &human7
// 	ischange_p(p_human7)
// 	fmt.Printf("human7: %v\n", human7)
// 	//可以发现在函数内部结构体已经发生了改变，而且还影响到了函数外部

// 	/*
// 		结构体的嵌套：
// 			go语言没有面向对象编程思想，也没有继承关系，但是可以通过结构体嵌套来实现这种效果

// 	*/

// 	type family struct {
// 		id     int
// 		name   string
// 		father Human
// 		mother Human
// 		kid    Human
// 	}

// 	var family1 = family{id: 1, name: "family1"}
// 	family1.father = human1
// 	family1.mother = human2
// 	family1.kid = human3
// 	fmt.Printf("family1: %v\n", family1)

// }

// func ischange_m(a Human) {
// 	a.id = 999
// 	fmt.Printf("a: %v\n", a)
// }
// func ischange_p(a *Human) {
// 	a.id = 9999
// 	fmt.Printf("a: %v\n", a)
// }
