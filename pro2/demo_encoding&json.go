// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"os"
// )

// type human struct {
// 	Name    string
// 	Age     int
// 	Email   string
// 	Parents []human
// }

// // !!!!!!!!!!!!!结构体内部属性名字一定要大写开头

// func test0() {
// 	p := human{
// 		Name:  "1",
// 		Age:   20,
// 		Email: "@@",
// 	}
// 	b, err := json.Marshal(p)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	fmt.Printf("b: %v\n", string(b))

// 	b = []byte(`{"Name":"1","Age":20,"Email":"@@"}`) //注意是反引号
// 	var p2 human
// 	json.Unmarshal(b, &p2)
// 	// Unmarshal解析json编码的数据，并将结果存储在v指向的值中。
// 	// 如果v为nil或不是指针，Unmarshal将返回一个InvalidUnmarshalError。
// 	fmt.Printf("p2: %v\n", p2)

// 	p3 := human{
// 		Name:    "3",
// 		Age:     18,
// 		Parents: []human{p, p2},
// 	}
// 	b2, _ := json.Marshal(p3)
// 	fmt.Printf("p3: %v\n", p3)
// 	fmt.Printf("b2: %v\n", string(b2))
// 	fmt.Printf("p3.Parents[1].Age: %v\n", p3.Parents[1].Age)

// 	b2 = []byte(`{"Name":"3","Age":18,"Email":"","Parents":[{"Name":"1","Age":20,"Email":"@@","Parents":null},{"Name":"1","Age":20,"Email":"@@","Parents":null}]}`)
// 	var p4 human
// 	json.Unmarshal(b2, &p4)
// 	fmt.Printf("p4: %v\n", p4)
// 	fmt.Printf("p4.Parents[1].Age: %v\n", p4.Parents[1].Age)

// }
// func main() {
// 	/*
// 		这个包可以实现json的编码和解码，就是将json字符串转换为struct或者将struct转换为json
// 		两个核心函数：
// 				1.json.Marshal()	//将struct编码成json，可以接受任意类型
// 					func json.Marshal(v interface{}) ([]byte, error)
// 				2.json.Unmarshal()	//将json转码成struct结构体
// 					func json.Unmarshal(data []byte, v interface{}) error

// 		两个核心结构体:
// 			1.type Decoder struct {	//从输入流读取并解析json
// 					r       io.Reader
// 					buf     []byte
// 					d       decodeState
// 					scanp   int   // start of unread data in buf
// 					scanned int64 // amount of data already scanned
// 					scan    scanner
// 					err     error

// 					tokenState int
// 					tokenStack []int
// 				}
// 			2.type Encoder struct {	//写json到输出流
// 					w          io.Writer
// 					err        error
// 					escapeHTML bool

// 					indentBuf    *bytes.Buffer
// 					indentPrefix string
// 					indentValue  string
// 				}
// 	*/
// 	lxf := human{
// 		Name:  "lxf",
// 		Age:   45,
// 		Email: "9@",
// 	}
// 	xh := human{
// 		Name:  "xh",
// 		Age:   42,
// 		Email: "0@",
// 	}
// 	xxh := human{
// 		Name:    "xxh",
// 		Age:     19,
// 		Email:   "1@",
// 		Parents: []human{lxf, xh},
// 	}

// 	b, _ := json.Marshal(xxh)
// 	fmt.Printf("b: %v\n", string(b))

// 	b1 := []byte(`{"Name":"xxh","Age":19,"Email":"1@","Parents":[{"Name":"lxf","Age":45,"Email":"9@","Parents":null},{"Name":"xh","Age":42,"Email":"0@","Parents":null}]}`)
// 	var m map[string]interface{}
// 	json.Unmarshal(b1, &m)
// 	for k, v := range m {
// 		fmt.Printf("k: %v,", k)
// 		fmt.Printf("v: %v\n", v)
// 	}

// }
// func test() {
// 	r, _ := os.Open("a.json")
// 	w, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 777)
// 	defer r.Close()
// 	defer w.Close()
// 	dec := json.NewDecoder(r)
// 	// NewDecoder返回一个从r读取的新解码器。
// 	enc := json.NewEncoder(w)
// 	// NewEncoder返回一个新的编码器，写入w。
// 	for {
// 		var v human
// 		err := dec.Decode(&v)
// 		// Decode从它的输入中读取下一个json编码的值，并将其存储在v指向的值中。
// 		if err != nil {
// 			fmt.Printf("err: %v\n", err)
// 			break
// 		}

// 		fmt.Printf("v: %v\n", v)

// 		b, _ := json.Marshal(&v)
// 		var k human
// 		json.Unmarshal(b, &k)
// 		io.WriteString(w, fmt.Sprintln(k))

// 		err = enc.Encode(&v)
// 		// Encode将v的JSON编码写入流，后跟换行符。
// 		if err != nil {
// 			fmt.Printf("err2: %v\n", err)
// 			break
// 		}
// 	}
// }
