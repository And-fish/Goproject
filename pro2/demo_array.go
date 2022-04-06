// package main

// import (
// 	"fmt"
// )

// func main() {
// 	/*
// 		数组是相同数据类型的一组数据的集合
// 		一旦定义长度不可修改
// 		可以通过下标(索引)来访问元素
// 		var name [size]type
// 		会填充size个0(string类型数组是空字符串,bool类型是false,float类型是0.000000)

// 	*/

// 	var list1 [2]int
// 	var list2 [3]string
// 	var list3 [2]bool
// 	var list4 [4]float32
// 	fmt.Printf("list1: %v,%T\n", list1, list1)
// 	fmt.Printf("list2: %v,%T\n", list2, list2)
// 	fmt.Printf("list3: %v,%T\n", list3, list3)
// 	fmt.Printf("list4: %f,%T\n", list4, list4)

// 	/*
// 		数组初始化var name = [size]type{value}
// 		如果value个数小于size，会按照定义时候的规则补齐
// 		[...]代表自适应value的个数
// 		{value}中可以通过索引指定位置初始化,{index1:value1,index5:value5},其他未指定的会自动补齐
// 	*/
// 	var a1 = [...]int{1, 2}
// 	a2 := [3]string{"a", "b", "c"}
// 	var a3 = [4]bool{true}
// 	a4 := [...]int{1: 5, 4: 6}
// 	fmt.Printf("a1: %v\n", a1)
// 	fmt.Printf("a2: %v\n", a2)
// 	fmt.Printf("a3: %v\n", a3)
// 	fmt.Printf("a4: %v\n", a4)

// 	/*
// 		通过索引可以修改数组中的值
// 	*/
// 	fmt.Printf("a2: %v\n", a2)
// 	a2[1] = "g"
// 	fmt.Printf("a2改: %v\n", a2)

// 	//遍历
// 	for i := 0; i < len(a2); i++ {
// 		fmt.Printf("%v ", a2[i])
// 	}
// 	fmt.Println("")
// 	for _, v := range a2 {
// 		fmt.Printf("%v ", v)
// 	}
// 	fmt.Println("")
// }