// package main

// import "fmt"

// const (
// 	default_Heat  = false
// 	default_Sugar = 5
// )

// type Coffee struct {
// 	Id    string
// 	Heat  bool
// 	Sugar int
// }
// type options struct {
// 	Heat  bool
// 	Sugar int
// }
// type setContainer interface {
// 	apply(*options)
// }
// type optionsFunc func(*options)

// func (f optionsFunc) apply(o *options) {
// 	f(o)
// }
// func WithHeat(h bool) setContainer {
// 	return optionsFunc(func(o *options) {
// 		o.Heat = h
// 	})
// }
// func WithSugar(s int) setContainer {
// 	return optionsFunc(func(o *options) {
// 		o.Sugar = s
// 	})
// }
// func get_newcoffee(id string, ops ...setContainer) *Coffee {
// 	opt := options{
// 		Heat:  default_Heat,
// 		Sugar: default_Sugar,
// 	}
// 	for _, v := range ops {
// 		v.apply(&opt)
// 	}
// 	return &Coffee{
// 		Id:    id,
// 		Heat:  opt.Heat,
// 		Sugar: opt.Sugar,
// 	}
// }
// func main() {
// 	/*
// 		先定义默认参数
// 		创建两个结构体:一个是用来承载和修改默认参数的，一个是用来返回最终信息的
// 		定义一个函数类型和接口类型，用来定义修改参数的函数apply()
// 		定义两个闭包函数，用来承载要修改的参数和修改的方法
// 		定义一个修改参数的函数apply()

// 		apply()表示，哪个闭包函数使用它，就执行哪个闭包函数。
// 		闭包函数的参数是要修改的值，使用这个闭包函数的时候，
// 			返回的东西即代表的定义的函数，又代表了内存中对应的参数。

// 		v.apply(&x)代表的是。因为v不仅代表了一个函数，还有内存中对应的一个参数y。
// 			所以这个apply函数就变成了，对y执行以x为参数的v

// 	*/
// 	c := get_newcoffee("01")
// 	fmt.Printf("c: %v\n", c)
// 	setsugar := WithSugar(8)
// 	c2 := get_newcoffee("02", setsugar)
// 	fmt.Printf("c2: %v\n", c2)
// 	setsugar = WithSugar(2)
// 	setheat := WithHeat(true)
// 	c3 := get_newcoffee("03", setsugar, setheat)
// 	fmt.Printf("c3: %v\n", c3)

// }
