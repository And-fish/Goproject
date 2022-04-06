// package main

// import "fmt"

// type Music interface {
// 	playmusic()
// }
// type Video interface {
// 	playvideo()
// }

// type Phone struct {
// 	id int
// }
// type Computer struct {
// 	id int
// }

// func (m Phone) playmusic() {
// 	fmt.Println("music...")
// }
// func (m Computer) playmusic() {
// 	fmt.Println("music...")
// }
// func (m Phone) playvideo() {
// 	fmt.Println("video...")
// }
// func (m Computer) playvideo() {
// 	fmt.Println("video...")
// }

// func main() {
// 	phone := Phone{1}
// 	phone.playmusic()
// 	phone.playvideo()

// 	computer := Computer{1}
// 	computer.playmusic()
// 	computer.playvideo()

// 	var musicplay Music
// 	musicplay = Phone{2}
// 	musicplay.playmusic()
// 	// musicplay.playvideo()报错，因为musicplay是用music接口定义的，接口中没有playmusic方法

// }
