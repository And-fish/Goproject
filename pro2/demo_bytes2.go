// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"os"
// )

// func btyes_reader() {
// 	data := "123456789"
// 	re := bytes.NewReader([]byte(data))
// 	fmt.Printf("re.Len(): %v\n", re.Len())
// 	fmt.Printf("re.Size(): %v\n", re.Size())

// 	buf := make([]byte, 2)
// 	for {
// 		n, err := re.Read(buf)
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println(string(buf[:n]))
// 	}
// 	fmt.Println("##########################################")
// 	re.Seek(0, 0)
// 	for {
// 		b, err := re.ReadByte()
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println(string(b))
// 	}
// 	fmt.Println("###########################")
// 	re.Seek(0, 0)
// 	off := int64(0)
// 	for {
// 		n, err := re.ReadAt(buf, off)
// 		if err != nil {
// 			break
// 		}
// 		off += int64(n)
// 		fmt.Println(string(buf[:n]))
// 	}
// }

// func bytes_testbuffer() {
// 	var b bytes.Buffer
// 	fmt.Printf("b: %v\n", b)
// 	var b1 = bytes.NewBufferString("hello")
// 	fmt.Printf("b1: %v\n", b1)
// 	var b2 = bytes.NewBuffer([]byte("hello"))
// 	fmt.Printf("b2: %v\n", b2)
// 	var b3 = new(bytes.Buffer)
// 	fmt.Printf("b3: %v\n", b3)
// 	// Buffer自带read和write方法
// 	// type Buffer struct {
// 	// 		buf      []byte
// 	// 		off      int
// 	// 		lastRead readOp
// 	// }
// 	// func NewBufferString(s string) *Buffer {
// 	// 	return &Buffer{buf: []byte(s)}
// 	// }
// }
// func bytes_buffer() {

// }
// func main() {
// 	/*
// 		bytes中的reader实现了
// 			io.Reader | io.ReaderAt | io.WriterTo | is.Seeker | io.ByteScanner | io.RuneScanner
// 		reader是只读的，可以seek
// 	*/
// 	var b bytes.Buffer
// 	n, _ := b.WriteString("hello")
// 	fmt.Println(n, string(b.Bytes()))

// 	b1 := make([]byte, 2)
// 	b.Read(b1)
// 	fmt.Printf("string(b1): %v\n", string(b1))
// 	file, _ := os.Open("a.txt")
// 	b.Reset()
// 	b.ReadFrom(file)
// 	fmt.Printf("b: %v\n", string(b.Bytes()))
// }
