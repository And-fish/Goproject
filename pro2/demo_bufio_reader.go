// package main

// import (
// 	"bufio"
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// func test1() {
// 	// r := strings.NewReader("get stringconst")
// 	r, _ := os.Open("a.txt")
// 	r2 := bufio.NewReader(r)
// 	s, _ := r2.ReadString('\n')
// 	fmt.Printf("s: %v\n", s)
// }
// func test2() {
// 	s := strings.NewReader("fucking world!!")
// 	str := strings.NewReader("happly newyear")
// 	msg := strings.NewReader("123456789100")
// 	br := bufio.NewReader(str)
// 	b, _ := br.ReadString('\n')
// 	fmt.Printf("b: %v\n", b)
// 	br.Reset(s)
// 	b, _ = br.ReadString('\n')
// 	fmt.Printf("b: %v\n", b)
// 	br = bufio.NewReaderSize(msg, 1)
// 	b, _ = br.ReadString('\n')
// 	fmt.Printf("b: %v\n", b)
// }
// func test3() {
// 	s := strings.NewReader("adhgsfasdklfhasioghasldkvcxvvcvb")
// 	br := bufio.NewReader(s)
// 	buf := make([]byte, 10)
// 	for {
// 		n, err := br.Read(buf)
// 		if err != nil {
// 			log.Fatal(err)
// 			break
// 		} else {
// 			fmt.Printf("string(buf[0:n]): %v\n", string(buf[:n]))
// 		}
// 	}
// }
// func test4() {
// 	s := strings.NewReader("akjfsiudfhsagdslajgla")
// 	br := bufio.NewReader(s)

// 	c, _ := br.ReadByte()
// 	fmt.Printf("c: %v\n", string(c))

// 	c, _ = br.ReadByte()
// 	fmt.Printf("c: %v\n", string(c))

// 	br.UnreadByte() //返回最后一个读取的字节
// 	c, _ = br.ReadByte()
// 	fmt.Printf("c: %v\n", string(c))
// }
// func test5() {
// 	a := strings.NewReader("gdhag\nLUIFI\nOSADUF")
// 	br := bufio.NewReader(a)

// 	// w, isprefix, _ := br.ReadLine()
// 	// fmt.Printf("w: %v,%v\n", string(w), isprefix)
// 	// // 如果一行的长度超过了缓冲大小，就会返回isprefix==true，而且超过的不会读取

// 	w, _ := br.ReadBytes('\n')
// 	fmt.Printf("w: %v\n", string(w))
// }
// func test6() {
// 	a := strings.NewReader("gdhag5,LUIFI,OSADUF")
// 	br := bufio.NewReader(a)

// 	w, _ := br.ReadSlice(',')
// 	fmt.Printf("string(w): %v\n", string(w))

// 	w, _ = br.ReadBytes(',')
// 	fmt.Printf("w: %v\n", string(w))
// }
// func test7() {
// 	a := strings.NewReader("gdhag5,LUIFI,OSADUF")
// 	br := bufio.NewReader(a)
// 	b := bytes.NewBuffer(make([]byte, 0))

// 	br.WriteTo(b)
// 	fmt.Printf("b: %v\n", b)
// }

// func main() {
// 	test7()
// }
