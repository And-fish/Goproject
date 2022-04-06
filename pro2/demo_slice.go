// package main

// import "fmt"

// func main() {
// 	/*
// 		切片和数组类似，但是数组是固定长度、不可修改的，但是切片长度是可修改的，拥有自动扩容的功能
// 		var name []type
// 		var name []type=make([]type,len)
// 		neme:=mkae([]type,len)
// 		make([]type,length,capacity)
// 		make中的length是初始长度，填补方式和数组一样
// 		capacity是容量,当长度>容量是,会扩容,
// 		??容量是用来开辟存储空间的??
// 								if长度<1024,容量=容量*2
// 								else 容量=容量*1.25
// 		##类似于python的列表
// 	*/
// 	var slice1 []int
// 	fmt.Printf("(slice1 == nil): %v\n", (slice1 == nil))
// 	fmt.Printf("slice1: %v\n", slice1)
// 	slice1 = append(slice1, 5, 6)
// 	fmt.Printf("slice1: %v\n", slice1)
// 	slice1 = append(slice1, slice1[0:1]...)
// 	fmt.Printf("slice1: %v\n", slice1)

// 	var slice2 = make([]int, 3)
// 	//var slice2 []int = make([]int, 0)
// 	fmt.Printf("slice2: %v\n", slice2)

// 	var slice3 = []int{1, 2, 3}
// 	fmt.Printf("slice3: %v\n", slice3)

// 	var slice4 = []int{1, 2, 3, 4}
// 	fmt.Printf("lenght:%v,capacity:%v\n", len(slice4), cap(slice4))
// 	slice4 = append(slice4, 5)
// 	fmt.Printf("lenght:%v,capacity:%v\n", len(slice4), cap(slice4))

// 	var array1 = [...]int{1, 2, 3, 4}
// 	array2 := array1
// 	fmt.Printf("array2: %T\n", array2) //[4]int数组类型
// 	slice5 := array1[:]
// 	fmt.Printf("slice5: %T\n", slice5) //[]int切片类型
// 	slice6 := array1[1:4]
// 	slice7 := slice5[1:4]
// 	var slice_copy = make([]int, 4)
// 	copy(slice_copy, slice5)
// 	fmt.Printf("slice_copy: %v\n", slice_copy)
// 	fmt.Printf("slice5: %v,slice6:%v,slice7:%v,slice_copy: %v\n", slice5, slice6, slice7, slice_copy)
// 	fmt.Printf("slice5: %v,slice6:%v,slice7:%v,slice_copy: %v\n", &slice5[2], &slice6[1], &slice7[1], &slice_copy[2])
// 	slice5[2] = 10
// 	fmt.Printf("slice5: %v,slice6:%v,slice7:%v,slice_copy: %v\n", slice5, slice6, slice7, slice_copy)
// 	/*
// 		slice5: [1 2 10 4],slice6:[2 10 4],slice7:[2 10 4]
// 		可以发现，直接用数组或者切片来赋值切片并不是很好的选择，
// 		因为数组和切片都是指向value所在的地址，修改对应位置的值意味着修改地址对应的值，会直接修改所有指向这个地址的切片和数组
// 		但是如果我们使用copy函数就能规避这个问题，copy函数会重新指向一个新的地址，所以改变原来地址对应的值并不会改变copy的切片
// 		copy(目的,源)，目的和源的类型要一致，copy不会修改目的的容量和长度，所以必须要用make声明长度和容量，如果长度小于源的长度，只会把前几个copy过去
// 		例如[1 2 3 4]copy到
// 							var slice []int，得到的slice是[]，因为默认len和cap都是0
// 							var slice =make([]int,4,5)，得到的是[1 2 3 4]，因为len正好等于4
// 							var silce =mkae([]int,5,5)，得到的是[1 2 3 4 0]，因为目的len>源len，所以只会修改前几个
// 							var slice =make([]int,3,4)，得到的是[1 2 3]，因为目的len<源len，也只会修改前几个
// 		所以，个人建议如果想得到和源一样的类型，应该使用var name=make([]type,len(源),cap(源))，copy(目的，源)
// 	*/

// 	var slice8 = []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	slen := len(slice8)
// 	for i := 0; i < slen; i++ {
// 		fmt.Printf("slice8[%v]: %v ", i, slice8[i])
// 	}
// 	fmt.Println("")

// 	i := 0
// 	for {
// 		if i == slen {
// 			break
// 		}
// 		fmt.Printf("slice8[%v]: %v ", i, slice8[i])
// 		i++
// 	}

// 	fmt.Println("")
// 	for i, v := range slice8 {
// 		fmt.Printf("slice8[%v]: %v ", i, v)
// 	}
// 	fmt.Println("")

// 	var slice9 []int
// 	slice9 = append(slice9, 5)
// 	slice9 = append(slice9, slice9[0]+5)
// 	slice9 = append(slice9, 15, 20, 25, 30, 35)
// 	fmt.Printf("slice9: %v\n", slice9)
// 	rmindex := 3
// 	slice9 = append(slice9[:rmindex], slice9[rmindex+1:]...)
// 	fmt.Printf("slice9改: %v\n", slice9)
// 	slice9 = append([]int{0}, slice9...)
// 	fmt.Printf("slice9再改: %v\n", slice9)
// 	slice9[0] = 1
// 	fmt.Printf("slice9: %v\n", slice9)
// 	key := 15
// 	for i, v := range slice9 {
// 		if key == v {
// 			fmt.Printf("i: %v\n", i)
// 			break
// 		}
// 	}

// }
