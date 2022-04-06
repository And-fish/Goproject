// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var str1 string = "hello world"
// 	var str2 string = `hello
// 	world
// 	!`
// 	fmt.Printf("str1: %v\n", str1)
// 	fmt.Printf("str2: %v\n", str2)

// 	var (
// 		s1 = "hello"
// 		s2 = "world"
// 	)
// 	fmt.Printf("s1 + s2: %v\n", s1+s2)

// 	msg := fmt.Sprintf("%s%s", s1, s2)
// 	fmt.Printf("msg: %v\n", msg)

// 	msg2 := strings.Join([]string{s1, s2}, "")
// 	fmt.Printf("msg2: %v\n", msg2)

// 	var msg3 bytes.Buffer
// 	msg3.WriteString(s1)
// 	msg3.WriteString(s2)
// 	fmt.Printf("msg3: %v\n", msg3.String())

// 	str3 := "hello world"
// 	fmt.Printf("str3[3]: %v\n", str3[3])
// 	fmt.Printf("str3: %c\n", str3[3])
// 	fmt.Printf("str3[3:5]: %v\n", str3[3:5])
// 	fmt.Printf("str3[3:]: %v\n", str3[3:])
// 	fmt.Printf("str3[:5]: %v\n", str3[:5])

// 	fmt.Printf("len(str3): %v\n", len(str3))
// 	fmt.Printf("strings.Split(str3, \" \"): %v\n", strings.Split(str3, " "))
// 	fmt.Printf("strings.Contains(str3, \"llo\"): %v\n", strings.Contains(str3, "llo"))
// 	fmt.Printf("strings.ToLower(str3): %v\n", strings.ToLower(str3))
// 	fmt.Printf("strings.ToUpper(str3): %v\n", strings.ToUpper(str3))
// 	fmt.Printf("strings.HasPrefix(str3, \"hello\"): %v\n", strings.HasPrefix(str3, "hello"))
// 	fmt.Printf("strings.HasSuffix(str3, \"hello\"): %v\n", strings.HasSuffix(str3, "hello"))
// 	fmt.Printf("strings.Index(str3, \"l\"): %v\n", strings.Index(str3, "l"))
// 	fmt.Printf("strings.LastIndex(str3, \"l\"): %v\n", strings.LastIndex(str3, "l"))

// 	a := "abc"
// 	b := "Abc"
// 	fmt.Printf("%q:%d\n", []byte(a), []byte(a))
// 	fmt.Printf("%q:%d\n", []byte(b), []byte(b))
// 	fmt.Println(strings.Compare(a, b)) //等于0则相等,a>b为1，a<b为-1
// 	fmt.Println(a == b)
// 	fmt.Print(strings.EqualFold(a, b))
// }
