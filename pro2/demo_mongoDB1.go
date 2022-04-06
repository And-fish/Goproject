// package main

// import (
// 	"context"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var client *mongo.Client

// type human struct {
// 	Name string
// 	Age  int
// }

// func init_DB() {
// 	co := options.Client().ApplyURI("mongodb://localhost:27017")
// 	var err error
// 	client, err = mongo.Connect(context.TODO(), co) //TODO返回一个非空的，空的Context。
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}

// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	fmt.Println("MongoDB连接成功")
// }
// func mongo_insertOne() {
// 	h1 := human{Name: "xxh", Age: 19}
// 	collection := client.Database("go_db").Collection("human")
// 	ior, err := collection.InsertOne(context.TODO(), h1)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("ior.InsertedID: %v\n", ior.InsertedID)
// 	}
// }
// func mongo_insertMany() {
// 	h1 := human{Name: "xxh", Age: 20}
// 	h2 := human{Name: "xzy", Age: 10}
// 	collection := client.Database("go_db").Collection("human")
// 	ior, err := collection.InsertMany(context.TODO(), []interface{}{h1, h2})
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("ior.InsertedID: %v\n", ior.InsertedIDs)
// 	}
// }
// func mongo_find() {
// 	collection := client.Database("go_db").Collection("human")
// 	c, err := collection.Find(context.TODO(), bson.D{})
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	defer c.Close(context.TODO())
// 	for c.Next(context.TODO()) {
// 		var result bson.D
// 		c.Decode(&result)
// 		fmt.Printf("result: %v\n", result)
// 		fmt.Printf("result.Map(): %v\n", result.Map())
// 	}
// }
// func mongo_updata() {
// 	collection := client.Database("go_db").Collection("human")
// 	update_data := bson.D{{"$set", bson.D{{"name", "xxh3"}, {"age", 20}}}}
// 	// ur, err := collection.UpdateOne(context.TODO(), bson.D{}, update_data)
// 	// // 只更新一个
// 	ur, err := collection.UpdateMany(context.TODO(), bson.D{}, update_data)
// 	// 更新所有匹配到的,参数二是用来匹配的,参数三是用来执行的
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("ur.UpsertedID: %v\n", ur.MatchedCount)     //匹配到的数量
// 		fmt.Printf("ur.ModifiedCount: %v\n", ur.ModifiedCount) //所修改的数量
// 	}
// }
// func mongo_del() {
// 	collection := client.Database("go_db").Collection("human")
// 	dr, err := collection.DeleteMany(context.TODO(), bson.D{{"name", "xxh"}})
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("dr.DeletedCount: %v\n", dr.DeletedCount)
// 	}
// }
// func main() {
// 	init_DB()
// 	defer client.Disconnect(context.TODO())
// 	// mongo_insertOne()
// 	// mongo_insertMany()
// 	// mongo_find()
// 	// mongo_updata()
// 	mongo_find()
// 	mongo_del()
// 	mongo_find()
// 	fmt.Println("end>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
// }
