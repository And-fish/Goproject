// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func createFile(name string) {
// 	f, err := os.Create(name)
// 	// func os.Create(name string) (*os.File, error)
// 	// 如果文件不存在就新建，如果文件存在就删掉原来的再新建
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("f.Name(): %v\n", f.Name())
// 	}
// }
// func createDir(name string) {
// 	// err := os.Mkdir(name, os.ModePerm)
// 	// // 创建单个
// 	// // func Mkdir(name string, perm FileMode) error
// 	// // 只返回错误反馈
// 	// if err != nil {
// 	// 	fmt.Printf("err: %v\n", err)
// 	// }

// 	err := os.MkdirAll(name, os.ModePerm)
// 	// func os.MkdirAll(path string, perm fs.FileMode) error
// 	// 只返回错误反馈
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}

// }
// func remove(name string) {
// 	// err := os.Remove(name)
// 	// // func os.Remove(name string) error
// 	// // 只能删除单个文件，不能删除有子文件或者子目录的目录
// 	// if err != nil {
// 	// 	fmt.Printf("err: %v\n", err)
// 	// }

// 	err := os.RemoveAll(name)
// 	// func os.RemoveAll(path string) error
// 	// RemoveAll删除路径及其包含的任何子路径
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }
// func getWD() {
// 	f, err := os.Getwd()
// 	// func os.Getwd() (dir string, err error)
// 	// Getwd返回与当前目录对应的根路径名。
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("f: %v\n", f)
// 	}
// }
// func chWD(name string) {
// 	err := os.Chdir(name)
// 	// func os.Chdir(dir string) error
// 	// Chdir将当前工作目录更改为指定目录。
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }
// func getTemp() {
// 	d := os.TempDir()
// 	// func os.TempDir() string
// 	// TempDir返回临时文件使用的默认目录。
// 	fmt.Printf("d: %v\n", d)
// }
// func renameFile(source string, target string) {
// 	err := os.Rename(source, target)
// 	// Rename(oldpath string, newpath string) error
// 	// Rename重命名(移动)旧路径到newpath。
// 	// 如果newpath已经存在并且不是一个目录，Rename将替换它。
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }

// func readFile(name string) {
// 	txt, err := os.ReadFile(name)
// 	// func os.ReadFile(name string) ([]byte, error)
// 	// ReadFile读取指定的文件并返回其内容。
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("txt: %v\n", string(txt[:]))
// 	}
// }

// func writeFile(name string, data string) {
// 	err := os.WriteFile(name, []byte(data), os.ModePerm)
// 	// func os.WriteFile(name string, data []byte, perm fs.FileMode) error
// 	// riteFile将数据写入指定的文件，并在必要时创建该文件
// 	// 会覆盖
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }
// func readDir(name string) {
// 	D, err := os.ReadDir(name)
// 	// func os.ReadDir(name string) ([]fs.DirEntry, error)
// 	// ReadDir读取指定的目录，返回所有按文件名排序的目录项。
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		for _, a := range D {
// 			fmt.Printf("a.Name(): %v\n", a.Name())
// 		}
// 	}

// }
// func main() {
// 	/*
// 		E:\GO\src\os
// 		os:
// 			提供给我们一个平台无关性的操作系统功能接口，采用类 Unix 设计，隐藏了不同操作系统间的差异，让不同的文件系统和操作系统对象表现一致。
// 	*/
// 	createFile("a.txt")
// 	createDir("tset//a")
// 	remove("tset")
// 	getWD()
// 	readFile("a.txt")
// 	writeFile("a.txt", "hello")
// 	readDir("c:\\Users\\DELL\\Desktop\\GOproject")

// }

