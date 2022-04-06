// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"strings"
// )

// func test1() {
// 	s := strings.NewReader("abc 123 asd")
// 	bs := bufio.NewScanner(s)
// 	bs.Split(bufio.ScanWords)
// 	bs.Split(bufio.ScanBytes)
// 	bs.Split(bufio.ScanLines)
// 	// ScanWords是一个Scanner的分割函数，它返回每个用空格分隔的文本单词，并删除周围的空格。
// 	for bs.Scan() {
// 		fmt.Println(bs.Text())
// 	}
// }
// func main() {
// 	test1()
// }
