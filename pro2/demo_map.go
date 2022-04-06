// package main

// import "fmt"

// func main() {
// 	/*
// 		{var name map[key_type]value_type
// 		name=make(map[ke_type]value_type)}
// 		name:=make(map[key_type]value_type)
// 	*/
// 	map1 := make(map[string]string)
// 	map1["name"] = "tom"
// 	map1["age"] = "19"
// 	map1["email"] = "1626245483@gmail.com"

// 	var map2 map[string]int
// 	map2 = make(map[string]int)
// 	map2["name"] = 5
// 	map2["age"] = 10

// 	var map3 = map[string]int{"name": 20, "age": 25}
// 	map4 := map[string]int{"name": 100, "age": 50}
// 	fmt.Printf("map1: %v\n", map1)
// 	fmt.Printf("map2: %v\n", map2)
// 	fmt.Printf("map3: %v\n", map3)
// 	fmt.Printf("map4: %v\n", map4)

// 	key := "name"
// 	value := map1[key]
// 	value2 := map1["age1"]
// 	v, isok := map1[key]
// 	v2, isok2 := map1["age1"]
// 	fmt.Printf("v: %v,isok:%v\n", v, isok)
// 	fmt.Printf("v2: %v,isok2:%v\n", v2, isok2)
// 	fmt.Printf("value: %v\n", value)
// 	fmt.Printf("value2: %v\n", value2)
// 	//如果存在返回value。如果不存在返回一个默认的空白(string为空格，int为0...)

// 	for key, value := range map1 {
// 		fmt.Printf("map1[%v]:%v ", key, value)
// 	}
// }
