// package main

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// func main() {
// 	fmt.Printf("os.Getpid(): %v\n", os.Getpid())   // 获取当前进程id
// 	fmt.Printf("os.Getppid(): %v\n", os.Getppid()) //获取父id

// 	attr := &os.ProcAttr{
// 		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
// 		Env:   os.Environ(),
// 		// Dir string 新进程的工作目录，如果配置了就会跳转目录
// 		// Env []string 新进程的环境变量列表
// 		// Files []*File 指定新进程继承的活动文件对象，对应标准输入，标准输出和标准错误输出。如果传入的是nil，file就是关闭的
// 	}

// 	p, err := os.StartProcess("C:\\Windows\\System32\\notepad.exe", []string{"C:\\Windows\\System32\\notepad.exe", "C:\\Users\\DELL\\Desktop\\GOproject\\pro2\\a.txt"}, attr)
// 	// StartProcess(name string, argv []string, attr *os.ProcAttr) (*os.Process, error)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	fmt.Printf("p: %v\n", p)
// 	fmt.Printf("p.Pid: %v\n", p.Pid)

// 	p2, _ := os.FindProcess(p.Pid) // 通过进程id查找进程
// 	fmt.Printf("p2: %v\n", p2)

// 	time.AfterFunc(time.Second*10, func() {
// 		p.Signal(os.Kill) // 向p发送kill指令
// 	})
// 	// 等待十秒执行函数

// 	ps, _ := p.Wait() //等待进程退出，然后返回
// 	fmt.Printf("ps.String(): %v\n", ps.String())
// }
