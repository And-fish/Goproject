// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"time"
// )

// func main() {
// 	// r := strings.NewReader("hello world")
// 	// func strings.NewReader(s string) *strings.Reader
// 	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 777)
// 	b, err := ioutil.ReadAll(f)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		// log.Fatal(err)
// 	}
// 	fmt.Printf("string(b): %v\n", string(b))

// 	d, _ := ioutil.ReadDir(".")
// 	for _, v := range d {
// 		fmt.Printf("v.Name(): %v\n", v.Name())
// 	}

// 	b, _ = ioutil.ReadFile("a.txt")
// 	fmt.Printf("string(b): %v\n", string(b))

// 	msg := []byte("fuck world!!!")
// 	err = ioutil.WriteFile("a.txt", msg, 755)
// 	// func ioutil.WriteFile(filename string, data []byte, perm fs.FileMode) error
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}

// 	tempfile()
// }

// func tempfile() {
// 	content := []byte("hello fucking world!!!")
// 	tmpfile, err := ioutil.TempFile("", "test")
// 	// 创建临时文件
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	fmt.Printf("tmpfile.Name(): %v\n", tmpfile.Name())
// 	// 查看临时文件所在位置
// 	defer os.Remove(tmpfile.Name())
// 	// defer，最后删除临时文件
// 	if _, err := tmpfile.Write(content); err != nil {
// 		// 向临时文件写入
// 		log.Fatal(err)
// 	}
// 	if err := tmpfile.Close(); err != nil {
// 		// 关闭临时文件
// 		log.Fatal(err)
// 	}

// 	attr := &os.ProcAttr{
// 		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
// 		Env:   os.Environ(),
// 		// Dir string 新进程的工作目录，如果配置了就会跳转目录
// 		// Env []string 新进程的环境变量列表
// 		// Files []*File 指定新进程继承的活动文件对象，对应标准输入，标准输出和标准错误输出。如果传入的是nil，file就是关闭的
// 	}
// 	p, err := os.StartProcess("C:\\Windows\\System32\\notepad.exe", []string{"C:\\Windows\\System32\\notepad.exe", tmpfile.Name()}, attr)
// 	time.AfterFunc(time.Second*10, func() { p.Signal(os.Kill) })
// 	p.Wait()
// }
