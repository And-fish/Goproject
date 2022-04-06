package main

import (
	"GOproject/pro3_pkg/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	service.TestUserservice()
	_ = gin.Default
}
