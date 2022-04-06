// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Printf("os.Environ(): %v\n", os.Environ())
// 	fmt.Printf("os.Getenv(\"GOPATH\"): %v\n", os.Getenv("GOPATH"))

// 	// func os.Setenv(key string, value string) error

// 	s, b := os.LookupEnv("GOPATH")
// 	fmt.Printf("b: %v\n", b)
// 	fmt.Printf("s: %v\n", s)

// 	fmt.Printf("os.Getenv(\"JAVA_HOME\"): %v\n", os.Getenv("JAVA_HOME"))
// 	fmt.Println(os.ExpandEnv("$JAVA_HOME>>>$GOPATH"))
// }
