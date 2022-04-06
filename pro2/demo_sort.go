// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type MY_uint []uint

// func (u MY_uint) Len() int           { return len(u) }
// func (u MY_uint) Less(i, j int) bool { return u[i] < u[j] }
// func (u MY_uint) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

// func main() {
// 	/*
// 		sort提供了排序切片和用户自定义数据集以及相关功能的函数
// 		sort包主要针对[]int、[]float64、[]string,以及其他自定义切片的排序

// 		type Interface interface {
// 			Len() int	//集合中的元素数量
// 			Less(i, j int) bool	//保证index==j之前的数是有序的
// 			Swap(i, j int)	//交换
// 		}

// 		type IntSlice []int
// 		type Float64Slice []float64
// 		type StringSlice []string

// 		func sort.Ints(x []int)
// 		// Ints按递增顺序对整型切片进行排序。
// 		func sort.IntsAreSorted(x []int) bool
// 		// IntsAreSorted报告切片x是否按递增顺序排序。
// 		func sort.SearchInts(a []int, x int) int
// 		// searchchints在一个已排序的整型切片中搜索x，并返回Search指定的索引。
// 		以上三种方法[]float64和[]string也有对应的函数

// 		Sort(data sort.Interface)	//排序
// 		Stable(data sort.Interface)		//稳定地对数据进行排序，同时保持相等元素的原始顺序。
// 		func sort.Reverse(data sort.Interface) sort.Interface	//Reverse返回数据的reverse结构。
// 		func sort.IsSorted(data sort.Interface) bool		//sSorted报告数据是否排序。
// 		Search(n int, f func(int) bool) int		//查找index值

// 		自定义依据排序，只需要自定义一个类型，在less()中修改成自己所需要的的方法
// 	*/
// 	i := []int{2, 1, 3, 4}
// 	sort.Ints(i)
// 	fmt.Printf("i: %v\n", i)

// 	u := []uint{1, 3, 2}
// 	sort.Sort(MY_uint(u))
// 	fmt.Println(u)

// 	map1 := map[string]int{"a": 4, "b": 5}
// 	map2 := map[string]int{"a": 2, "b": 5}
// 	map3 := map[string]int{"a": 5, "b": 5}
// 	mp := mapSlice{map1, map2, map3}
// 	sort.Sort(mp)
// 	fmt.Printf("mp: %v\n", mp)

// 	d_int1 := [][]int{[]int{1, 4}, []int{9, 3}, []int{7, 5}}
// 	// d_int1:=[][]int{{1,4},{9,3},{7,5}}
// 	d := doubleInts(d_int1)
// 	sort.Sort(d)
// 	fmt.Printf("d: %v\n", d)

// 	human1 := human{50, "h1"}
// 	human2 := human{25, "h2"}
// 	human3 := human{75, "h3"}
// 	h := humanSlice{human1, human2, human3}
// 	fmt.Printf("h: %v\n", h)
// 	sort.Sort(h)
// 	fmt.Printf("h: %v\n", h)
// }

// type doubleInts [][]int

// //设置为根据index==1的大小来排序
// func (d doubleInts) Less(i, j int) bool { return d[i][1] < d[j][1] }
// func (d doubleInts) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
// func (d doubleInts) Len() int           { return len(d) }

// type mapSlice []map[string]int

// // 设置为依据map中key=="a"的值排序
// func (m mapSlice) Less(i, j int) bool { return m[i]["a"] < m[j]["a"] }
// func (m mapSlice) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
// func (m mapSlice) Len() int           { return len(m) }

// type human struct {
// 	id   int
// 	name string
// }
// type humanSlice []human

// // 设置为依据human的id排序
// func (m humanSlice) Less(i, j int) bool { return m[i].id < m[j].id }
// func (m humanSlice) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
// func (m humanSlice) Len() int           { return len(m) }
