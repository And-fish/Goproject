// package main

// import (
// 	"bufio"
// 	"bytes"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func test1() {
// 	f, _ := os.OpenFile("a.txt", os.O_RDWR, 777)
// 	defer f.Close()
// 	// f := bytes.NewBuffer(make([]byte, 0))
// 	bw := bufio.NewWriter(f)
// 	bw.WriteString("hahhahahaahha")
// 	bw.Flush() //要刷新一下，写入数据到底层的io.Writer
// 	// fmt.Printf("f: %v\n", f)

// }
// func test2() {
// 	b := bytes.NewBuffer(make([]byte, 0))
// 	bw := bufio.NewWriter(b)
// 	bw.WriteString("123456789")
// 	c := bytes.NewBuffer(make([]byte, 0))
// 	bw.Reset(c)
// 	bw.WriteString("abc")
// 	bw.Flush()
// 	fmt.Printf("b: %v\n", b)
// 	fmt.Printf("c: %v\n", c)
// }
// func test3() {
// 	b := bytes.NewBuffer(make([]byte, 0))
// 	bw := bufio.NewWriter(b)

// 	f, _ := os.OpenFile("a.txt", os.O_RDWR, 777)
// 	// f:=strings.NewReader("hjakfiusfaa")
// 	bw.ReadFrom(f)
// 	bw.Flush()
// 	fmt.Printf("b: %v\n", b)
// }

// func Reader_Writer() {
// 	b1 := strings.NewReader("1234567")
// 	b2 := bytes.NewBuffer(make([]byte, 0))
// 	br := bufio.NewReader(b1)
// 	bw := bufio.NewWriter(b2)
// 	rw := bufio.NewReadWriter(br, bw)

// 	r, _ := rw.ReadString('\n')
// 	fmt.Printf("string(r): %v\n", string(r))

// 	rw.WriteString("gagga")
// 	rw.Flush()
// 	fmt.Printf("b2: %v\n", b2)

// }
// func main() {
// 	/*
// 		bufio的writer功能和reader差不多，所以很多不会详细介绍
// 	*/
// 	Reader_Writer()
// }
