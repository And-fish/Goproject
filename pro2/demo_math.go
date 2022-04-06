// package main

// import (
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"time"
// )

// func math_const() {
// 	fmt.Printf("math.MaxInt: %v\n", math.MaxInt)
// 	fmt.Printf("math.MinInt: %v\n", math.MinInt)
// 	fmt.Printf("math.MaxFloat64: %v\n", math.MaxFloat64)
// 	fmt.Printf("math.SmallestNonzeroFloat64: %v\n", math.SmallestNonzeroFloat64)
// 	fmt.Printf("math.Pi: %v\n", math.Pi)
// }
// func math_fanction() {
// 	/* 	取绝对值，注意math中函数的传参和返回值都是float64类型的
// 	func math.Abs(x float64) float64 */
// 	fmt.Printf("math.Abs(f): %v\n", math.Abs(-3.85))

// 	/* 	返回x的y次方
// 	   	func math.Pow(x float64, y float64) float64 */
// 	fmt.Printf("math.Pow(2, 3): %v\n", math.Pow(2, 3))

// 	/* 	返回10的n次方
// 	   	func math.Pow10(n int) float64 */
// 	fmt.Printf("math.Pow10(2): %v\n", math.Pow10(2))

// 	/* 	返回x的开平方
// 	   	func math.Sqrt(x float64) float64 */
// 	fmt.Printf("math.Sqrt(9): %v\n", math.Sqrt(9))

// 	/* 	返回x的开立方
// 	   	func math.Cbrt(x float64) float64 */
// 	fmt.Printf("math.Cbrt(8): %v\n", math.Cbrt(8))

// 	/*	向上取整
// 		func math.Cbrt(x float64) float64 */
// 	fmt.Printf("math.Ceil(1.1): %v\n", math.Ceil(1.1))

// 	/* 	向下取整
// 	   	func math.Floor(x float64) float64 */
// 	fmt.Printf("math.Floor(1.8): %v\n", math.Floor(1.8))

// 	/* 	返回x对y的取余==x%y
// 	   	func math.Mod(x float64, y float64) float64 */
// 	fmt.Printf("math.Mod(10, 3): %v\n", math.Mod(10, 3))

// 	/* 	返回f的整数部分和小数部分
// 	   	func math.Modf(f float64) (int float64, frac float64) */
// 	i, f := math.Modf(3.1415)
// 	fmt.Printf("i: %v f: %v\n", i, f)
// 	fmt.Printf("i: %v f: %.5v\n", i, f)
// }
// func test_rand() {
// 	// UnixNano返回t作为Unix时间，即自UTC 1970年1月1日以来经过的纳秒数。
// 	rand.Seed(time.Now().UnixNano()) //设置一个种子，用time.Now().UnixNano()，保证输入的种子一定每次都是不同的
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("rand.Int(): %v\n", rand.Int())
// 	}
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("rand.Intn(100): %v\n", rand.Intn(100))
// 	}

// }
// func main() {
// 	/*
// 		该包包含了一些常量和一些有用的数学计算函数，例如:三角函数、随机数、绝对值、平方根等

// 	*/
// 	test_rand()

// }
