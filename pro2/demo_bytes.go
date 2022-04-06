// package main

// import (
// 	"bytes"
// 	"fmt"
// )

// func test() {
// 	var i int = 100
// 	var b byte = 10
// 	var s string = "hello world"
// 	b = byte(i)
// 	fmt.Printf("b: %v\n", b)
// 	s = string(b)
// 	fmt.Printf("s: %v\n", s)
// }
// func function() {
// 	s := "hello world"
// 	b := []byte(s)
// 	b1 := []byte("hello world !!!")
// 	b2 := []byte("HELLO world !!!")

// 	iscontains := bytes.Contains(b, b1)
// 	fmt.Printf("iscontains: %v\n", iscontains)
// 	iscontains = bytes.Contains(b1, b)
// 	fmt.Printf("iscontains: %v\n", iscontains)
// 	iscontains = bytes.Contains(b2, b)
// 	fmt.Printf("iscontains: %v\n", iscontains)
// 	// func bytes.Contains(b []byte, subslice []byte) bool
// 	// 查看前一个里面有没有包含后一个(会区分大小写)
// 	// 就像 strings.Contains("hello", "he")==True

// 	showcount := bytes.Count(b1, []byte("o"))
// 	fmt.Printf("showcount: %v\n", showcount)
// 	showcount = bytes.Count(b2, []byte("o"))
// 	fmt.Printf("showcount: %v\n", showcount)
// 	// func bytes.Count(s []byte, sep []byte) int
// 	// 查看参数1中有多少参数2(会区分大小写)
// 	// 就像 strings.Count("hooO","o")==2

// 	showRepeat := bytes.Repeat(b, 0)
// 	fmt.Printf("showRepeat: %v\n", string(showRepeat))
// 	showRepeat = bytes.Repeat(b, 2)
// 	fmt.Printf("showRepeat: %v\n", string(showRepeat))
// 	// func bytes.Repeat(b []byte, count int) []byte
// 	// 重复参数1[]byte，count次
// 	// 就像 strings.Repeat("ha",2)=="haha"

// 	replaceByte := []byte("abcb")
// 	showReplace := bytes.Replace(replaceByte, []byte("b"), []byte("p"), 1)
// 	fmt.Printf("replaceByte: %v\n", string(replaceByte))
// 	fmt.Printf("showReplace: %v\n", string(showReplace))
// 	showReplace = bytes.Replace(replaceByte, []byte("b"), []byte("p"), 100)
// 	fmt.Printf("showReplace: %v\n", string(showReplace))
// 	showReplace = bytes.Replace(replaceByte, []byte("b"), []byte("p"), -1)
// 	fmt.Printf("showReplace: %v\n", string(showReplace))
// 	showReplace = bytes.Replace(replaceByte, []byte("5"), []byte("p"), 100)
// 	fmt.Printf("showReplace: %v\n", string(showReplace))
// 	showReplace = bytes.ReplaceAll(replaceByte, []byte("b"), []byte("p"))
// 	fmt.Printf("showReplace: %v\n", string(showReplace))
// 	// func bytes.Replace(s []byte, old []byte, new []byte, n int) []byte
// 	// 把s []byte中的前n个old换成new，
// 	// 如果s中的old数量小于n个，不会报错。
// 	// 如果s中没有old，不会报错
// 	// 不会修改源byte，会返回一个修改过的byte
// 	// n==-1时，会修改所有的old
// 	// func bytes.ReplaceAll(s []byte, old []byte, new []byte) []byte
// 	// 把s []byte中所有的old都换成new，不会修改源byte，会返回修改过的byte
// 	// 就像strings.Replace("hello", "l", "a", 1)=="healo"
// 	// 就像strings.ReplaceAll("hello", "l", "a")=="heaao"

// 	runesByte := []byte("你好")
// 	runesByte2 := []byte("hello")
// 	showRunes := bytes.Runes(runesByte)
// 	showRunes2 := bytes.Runes(runesByte2)
// 	fmt.Printf("len(runesByte): %v\n", len(runesByte))
// 	fmt.Printf("len(showRunes): %v\n", len(showRunes))
// 	fmt.Printf("len(runesByte): %v\n", len(runesByte2))
// 	fmt.Printf("len(showRunes): %v\n", len(showRunes2))
// 	fmt.Printf("showRunes: %v\n", showRunes)
// 	fmt.Printf("showRunes2: %v\n", showRunes2)
// 	// func bytes.Runes(s []byte) []rune
// 	// 将byte从UTF-8转化为Unicode类型

// 	joinByte := [][]byte{[]byte("hello"), []byte("fucking"), []byte("world")}
// 	showJoin := bytes.Join(joinByte, []byte(","))
// 	fmt.Printf("showJoin: %v\n", string(showJoin))
// 	// Join(s [][]byte, sep []byte) []byte
// 	// 把二维byte用sep连接起来，返回一个[]byte
// 	// 就像strings.Join([]string{"hello","world"},",")
// }
// func main() {
// 	/*
// 		bytes包提供了对字节切片进行读写操作的一系列函数，
// 		字节切片处理的函数比较多分为基本处理函数、比较函数、后缀检查函数、索引函数、分割函数、大小写处理函数和子切片处理函数等
// 	*/
// 	function()
// }
