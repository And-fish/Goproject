// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// func openclose(name string) {
// 	// f, err := os.Open(name)
// 	// // return OpenFile(name, O_RDONLY, 0)
// 	// if err != nil {
// 	// 	fmt.Printf("err: %v\n", err)
// 	// } else {
// 	// 	fmt.Printf("f.Name(): %v\n", f.Name())
// 	// }

// 	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 755)
// 	/*
// 		至少要有O_RDONLY, O_WRONLY, or O_RDWR 中的一个
// 		O_RDONLY int = syscall.O_RDONLY // 只读
// 		O_WRONLY int = syscall.O_WRONLY // 只写
// 		O_RDWR   int = syscall.O_RDWR   // 读写
// 		下面的可以添加
// 		O_APPEND int = syscall.O_APPEND // 追加写入
// 		O_CREATE int = syscall.O_CREAT  // 不存在就创建
// 		O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
// 		O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
// 		O_TRUNC  int = syscall.O_TRUNC  // 如果存在会重新创建
// 	*/
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("f.Name(): %v\n", f.Name())
// 	}
// }

// func createFile(name string) {
// 	f, err := os.Create(name)
// 	// return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("f.Name(): %v\n", f.Name())
// 	}

// 	// f2,err:=os.CreateTemp("","Temp")
// 	// // 第一个参数是在指定目录下创建，如果没指定就默认在临时目录下创建
// 	// // 第二个参数是创建临时文件的前缀
// 	// // 最后得到的文件名是前缀+一串随机数字
// 	// if err != nil {
// 	// 	fmt.Printf("err: %v\n", err)
// 	// }else {
// 	// 	fmt.Printf("f2.Name(): %v\n", f2.Name())
// 	// }
// }
// func readFile(name string) {
// 	// f, _ := os.Open(name)
// 	// for {
// 	// 	buf := make([]byte, 10)
// 	// 	n, err := f.Read(buf)
// 	// 	// func (*os.File).Read(b []byte) (n int, err error)
// 	// 	// 文件结束，return 0,io.EOF
// 	// 	if err == io.EOF {
// 	// 		f.Close()
// 	// 		break
// 	// 	}
// 	// 	fmt.Printf("n: %v\n", n)
// 	// 	fmt.Printf("string(buf): %v\n", string(buf))
// 	// }

// 	// f, _ := os.Open(name)
// 	// buf := make([]byte, 4)
// 	// n, err := f.ReadAt(buf, 2)
// 	// // 从偏移值开始read，我猜不会移动偏移的位置(即如果再接着read会从头开始)
// 	// if err != io.EOF {
// 	// 	fmt.Printf("n: %v\n", n)
// 	// 	fmt.Printf("string(buf): %v\n", string(buf))
// 	// }
// 	// buf = make([]byte, 3)
// 	// n, err = f.Read(buf)
// 	// fmt.Printf("err: %v\n", err)
// 	// fmt.Printf("n: %v\n", n)
// 	// fmt.Printf("string(buf): %v\n", string(buf))

// 	f, _ := os.Open(name)
// 	for {
// 		f.Seek(0, 1)
// 		// 第一个参数表示偏移量，第二个参数表示参考位置（0是原点，1是当前位置，2是末尾）
// 		buf := make([]byte, 4)
// 		n, err := f.Read(buf)
// 		if err == io.EOF {
// 			f.Close()
// 			return
// 		}
// 		fmt.Printf("n: %v\n", n)
// 		fmt.Printf("string(buf): %v\n", string(buf))
// 	}
// 	f.Close()
// }
// func write(name string, data string) {
// 	f, _ := os.OpenFile(name, os.O_RDWR, 777)
// 	f.Write([]byte(data))
// 	// f.WriteString(data)
// 	// f.WriteAt([]byte(data), 3)
// 	// 偏转写入，不要使用追加模式
// 	f.Close()
// }

// func main() {
// 	// readFile("a.txt")
// 	// err := d_tree("c:\\Users\\DELL\\Desktop\\GOproject")
// 	// fmt.Printf("err: %v\n", err)
// 	write("a.txt", "1")
// }

// func d_tree(name string) error {
// 	_, errfile := os.ReadFile(name)
// 	_, errdir := os.ReadDir(name)
// 	switch {
// 	case errfile != nil && errdir == nil:
// 		msg := strings.Split(name, "\\")
// 		s := strings.Join(msg[len(msg)-1:], "")
// 		fmt.Printf("%v\n", s)
// 		return treeDir(name, 1)
// 	case errdir != nil && errfile != nil:
// 		return errdir
// 	default:
// 		msg := strings.Split(name, "\\")
// 		s := strings.Join(msg[len(msg)-1:], "")
// 		fmt.Printf("%v\n", s)
// 		return fmt.Errorf("目标是文件")
// 	}
// }
// func treeDir(name string, count int) error {
// 	D, err := os.ReadDir(name)
// 	if err != nil {
// 		return fmt.Errorf(err.Error())
// 	} else {
// 		for _, a := range D {
// 			for i := 0; i < count; i++ {
// 				fmt.Print("    ")
// 			}
// 			fmt.Printf("%v\n", a.Name())
// 			if a.IsDir() {
// 				treeDir(name+"\\"+a.Name(), count+1)
// 			}
// 		}
// 	}
// 	return nil
// }
