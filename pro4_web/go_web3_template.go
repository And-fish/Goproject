// package main

// import (
// 	"html/template"
// 	"net/http"
// 	"os"
// )

// func test0() {
// 	name := "熊玺皓"
// 	muban := "hello,{{.}}"
// 	t, _ := template.New("test").Parse(muban)
// 	t.Execute(os.Stdout, name)
// }
// func test1() {
// 	type Human struct {
// 		Name string
// 		Age  int
// 	}
// 	human1 := Human{Name: "熊玺皓", Age: 20}
// 	muban := "我是{{.Name}},我{{.Age}}岁"
// 	t, _ := template.New("test").Parse(muban)
// 	t.Execute(os.Stdout, human1)
// }
// func main() {
// 	/*
// 		templates定义了数据驱动的文本输出。生成html文件的模板在html/template包下面。
// 		模板使用插值语法{{.var}}格式，也可以使用一些流程控制。
// 		例如判断if else、循环range还可以使用一些函数，包括内建函数和自定义函数。

// 		test.html测试内容:
// 			测试hello wolde
// 			"output"
// 			output
// 			output
// 			output
// 			23 < 58
// 			23< 58
// 			23 <58
// 			23<58
// 			wude
// 			你还没成年

// 		test2.html测试内容:
// 			我很好!

// 		test2.html测试内容:
// 			这是header部分
// 			正文
// 			这是footer部分
// 	*/
// 	server()
// }
// func server() {
// 	// tmp1 := func(w http.ResponseWriter, r *http.Request) {
// 	// 	t, err := template.ParseFiles("./test.html")
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// 	t.Execute(w, "hello wolde")
// 	// }

// 	tmp3 := func(w http.ResponseWriter, r *http.Request) {
// 		t, err := template.ParseFiles("./test3.html", "./test.html", "./test2.html")
// 		if err != nil {
// 			panic(err)
// 		}
// 		t.Execute(w, nil)
// 	}

// 	// tmp2 := func(w http.ResponseWriter, r *http.Request) {
// 	// 	t, err := template.ParseFiles("./test2.html")
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// 	slice := []string{"我", "很", "好", "!"}
// 	// 	t.Execute(w, slice)
// 	// }

// 	server := http.Server{
// 		Addr: "127.0.0.1:9999",
// 	}
// 	// http.HandleFunc("/hello", tmp1)
// 	http.HandleFunc("/hello", tmp3)
// 	server.ListenAndServe()
// }
