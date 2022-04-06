// package main

// import (
// 	"context"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var client *mongo.Client

// func init_DB() {
// 	co := options.Client().ApplyURI("mongodb://localhost:27017")
// 	n, err := mongo.Connect(context.TODO(), co)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}

// 	err = n.Ping(context.TODO(), nil)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	fmt.Println("MongoDB连接成功")
// }
// func main() {
// 	init_DB()

// }
