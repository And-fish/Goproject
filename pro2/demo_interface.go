// package main

// import (
// 	"fmt"
// )

// type USB interface {
// 	read() string
// 	write() string
// 	change_p()
// 	change_v()
// }
// type Computer struct {
// 	id   int
// 	name string
// }
// type Phone struct {
// 	id   int
// 	name string
// }

// func (computer Computer) read() string {
// 	fmt.Println("reading...")
// 	return computer.name
// }
// func (computer Computer) write() string {
// 	fmt.Println("write...")
// 	return computer.name
// }
// func (phone Phone) read() string {
// 	fmt.Println("reading...")
// 	return phone.name
// }

// func (phone Phone) write() string {
// 	fmt.Println("write...")
// 	return phone.name
// }

// func (p *Phone) change_p() {
// 	p.name = "changed_p"
// }
// func (p *Computer) change_p() {
// 	p.name = "changed_p"
// }
// func (p Phone) change_v() {
// 	p.name = "changed_v"
// }
// func (p Computer) change_v() {
// 	p.name = "changed_v"
// }
// func main() {

// 	/*
// 		接口像是一个公司里面的领导，他会定义一些通用规范，只设计规范，不实现规范
// 		go语言的接口是一种新的类型定义，他把所有具有共性的方法定义在一起。任何其他类型只要实现了这些方法就是实现了这个接口
// 		语法格式和方法非常相似
// 		type name interface{
// 			method_name1 {return_type}
// 			method_name2 {return_type}
// 		}
// 		type struct_name struct{

// 		}
// 		func (struct_name_variable struct_name) method_name1() [return_type]{

// 		}
// 		func (struct_name_variable struct_name) method_name2() [return_type]{

// 		}
// 	*/
// 	computer1 := Computer{1, "computer1"}
// 	fmt.Printf("computer1.read(): %v\n", computer1.read())
// 	phone1 := Phone{1, "phone1"}
// 	fmt.Printf("phone1.read(): %v\n", phone1.read())

// 	m := Computer{2, "comperter"}
// 	fmt.Printf("m.name: %v\n", m.name)
// 	m.change_v()
// 	fmt.Printf("m.name: %v\n", m.name)
// 	m.change_p()
// 	fmt.Printf("m.name: %v\n", m.name)

// }
