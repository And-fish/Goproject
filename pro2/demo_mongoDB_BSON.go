// package main

// import (
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// func main() {
// 	/*
// 		MongoDB中的JSON文档储存在名为BSON(二进制编码的JSON)的二进制表示中
// 		BSON编码扩展了JSON表示，使其包含额外的类型，如int、long、date、浮点数和decimal128。
// 		使程序更容易可靠的处理、排序和比较数据

// 		连接MongoDB的GO驱动程序中有两大类型表示BSON数据: D和Raw

// 		类型D交租被用来简洁的构建使用本地go类型的BSON对象。D家族包括四类：
// 			1.D:一个BSON文档。这种类型应该在顺序重要的情况下使用，比如MongDB命令、
// 			2.M:一张无序的map。它和D是一样的，只是它不保持顺序
// 			3.A:一个BSON数组
// 			4.E:D里面的一个元素

// 		示例:使用D类型构建的过滤器文档的例子，它可以用来查找name字段与“张三”、“李四”匹配的文档
// 		bson.D{{
// 			"name",
// 			bson.D{{
// 				"$in",
// 				bson.A{"张三，李四"},
// 			}},
// 		}}

// 		Raw类型家族用于字节切片。还可以使用lookup()从原始类型检索单个元素
// 	*/

// 	//filter 过滤器 ==sql where
// 	d := bson.D{{"name", "xxh"}}
// 	fmt.Printf("d: %v\n", d)
// }
