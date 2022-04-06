// package main

// func main() {
// 	/*
// 		包可以区分命令空间(一个文件夹不能有两个同名文件)，可以方便去管理项目
// 		1.一个文件夹下面只能有一个package
// 			import后面的是GOPATH下面的相对路径(现在不一定)
// 		2.比如说实现了一个计算器的package，名叫“calc”位于calc下；
// 			但是想给别人一个使用范例，可以在calc下面建立一个example子目录，这个子目录里面有一个example.go文件。
// 			此时example.go可以使main包，里面还可以有个main函数
// 		3.一个package的文件不能在多个文件夹下
// 			如果多个文件夹下有重名的package，他们其实是彼此无关的package
// 			如果一个go文件需要同时使用多个目录下的同名package，需要在import这些目录时为每个目录指定一个别名
// 	*/
// }
