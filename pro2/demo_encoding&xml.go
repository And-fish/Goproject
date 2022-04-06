// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// type human struct {
// 	XMLNAME xml.Name `xml:"Human"`
// 	Name    string   `xml:"name"`
// 	Age     int      `xml:"age"`
// }

// func test() {
// 	p := human{
// 		Name: "xxh",
// 		Age:  19,
// 	}
// 	b, _ := xml.Marshal(p)
// 	// b, _ = xml.MarshalIndent(p, "%", "&") //有缩进格式，前一个是前缀，后一个是缩进标志
// 	// func xml.MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error)
// 	fmt.Printf("string(b): \n%v\n", string(b))

// 	s := `<human>
//     <Human></Human>
//     <name>xxh</name>
//     <age>19</age>
// </human>`
// 	b = []byte(s)
// 	var p2 human
// 	fmt.Printf("b: %v\n", string(b))
// 	xml.Unmarshal(b, &p2)
// 	fmt.Printf("p2: %v\n", p2)
// }
// func main() {
// 	/*
// 				xml包实现xml解析
// 				两个核心函数:
// 					1.xml.Marshal()	//将struct编码成xml，可以接受任意类型
// 						func xml.Marshal(v interface{}) ([]byte, error)
// 					2.xml.Unmarshal()	//将xml转码成struct结构体
// 						func xml.Unmarshal(data []byte, v interface{}) error
// 				两个核心结构体:
// 					1.type Decoder struct {	//从输入流读取并解析xml
// 						...
// 					}
// 					2.type Encoder struct {	//
// 		    			p printer
// 					}

// 	*/
// 	test1()
// }
// func test1() {
// 	b, _ := ioutil.ReadFile("a.xml")
// 	var p human
// 	xml.Unmarshal(b, &p)
// 	fmt.Printf("p: %v\n", p)

// 	r, _ := os.Open("a.xml")
// 	w, _ := os.OpenFile("b.xml", os.O_RDWR|os.O_APPEND, 777)
// 	defer w.Close()
// 	defer r.Close()

// 	var p2 human
// 	d := xml.NewDecoder(r)
// 	d.Decode(&p2)
// 	fmt.Printf("p2: %v\n", p2)
// 	e := xml.NewEncoder(w)
// 	e.Encode(&p)
// }
