// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// func testcopy() {
// 	r := strings.NewReader("hello world")
// 	writer, _ := os.OpenFile("a.txt", os.O_RDWR, 755)
// 	_, err := io.Copy(writer, r)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }
// func testcopyBuffer() {
// 	r := strings.NewReader("123456789")
// 	buf := make([]byte, 5)
// 	writer, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 755)
// 	_, err := io.CopyBuffer(writer, r, buf)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }
// func main() {
// 	/*
// 		io==input output
// 		IO操作被封装在了如下几个包中
// 		io为IO原语提供了基本的接口
// 		io/ioutil封装一些实用的I/O函数
// 		fmt实现了格式化I/O，类似于C语言的printf和scanf (format)
// 		bufio实现了带缓冲I/O
// 	*/

// 	/*
// 		在io包中最重要的两个接口就是read和write接口，只要实现了这两个接口他就有了I/O功能
// 		type Reader interface {
// 			Read(p []byte) (n int, err error)
// 		}
// 		type Writer interface {
// 			Write(p []byte) (n int, err error)
// 		}
// 	*/

// 	/*
// 		哪些类型实现了reader和writer接口？
// 			os.File:	reader\writer
// 			bytes.Buffer:	reader\writer
// 			crypto.Conn:	reader\writer
// 			tls.Conn:	reader\writer
// 			strings.Reader
// 			bufio.Reader
// 			bufio.Writer
// 			bytes.Reader
// 			gzip.Reader
// 			gzip.Writer
// 			compress.Reader
// 			compress.Writer
// 			crypto.StreamReader
// 			crypto.StreamWriter
// 			cipher.StreamReader
// 			cipher.StreamWriter
// 			encoding.Reader
// 			encoding.Writer
// 			csv.Reader
// 			csv.Writer
// 	*/
// 	testcopyBuffer()
// }
