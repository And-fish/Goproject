// package main

// import (
// 	"fmt"
// 	"time"
// )

// func showtime() {
// 	now := time.Now()
// 	fmt.Printf("now: %v\n", now)
// 	year := time.Now().Year()
// 	month := time.Now().Month()
// 	day := time.Now().Day()
// 	hour := time.Now().Hour()
// 	minute := time.Now().Minute()
// 	second := time.Now().Second()
// 	fmt.Printf("%v-%02d-%02v-%02v-%02v-%02v\n", year, month, day, hour, minute, second)
// 	fmt.Printf("%T,%T,%T,%T,%T,%T,%T\n", now, year, month, day, hour, minute, second)
// 	// 打印显示时间，其中now是time.Time类型，Month是time.Month类型
// }
// func timestamp() {
// 	// 时间戳是自1970年1月1日(08:00:00GMT)至当前时间的总毫秒数，
// 	// 也被称为Unix时间戳(UnixTimestamp)
// 	now := time.Now()
// 	fmt.Printf("now.Unix(): %T，%v\n", now.Unix(), now.Unix())
// 	// 除此之外还有纳秒时间戳，可以使用time().Now().UnixNano()
// 	fmt.Printf("now.UnixNano(): %T，%v\n", now.UnixNano(), now.UnixNano())
// 	// 通过时间戳获取时间
// 	timestamp := now.Unix()
// 	timeOBJ := time.Unix(timestamp, 0) //参数1是时间戳，参数2是纳秒
// 	fmt.Printf("timeOBJ: %v\n", timeOBJ)
// }
// func timeFunction() {
// 	now := time.Now()
// 	fmt.Printf("now: %v\n", now)
// 	fmt.Printf("changedTime: %v\n", now.Add(time.Hour*3+time.Minute*2+time.Second*50))
// 	//time.Hour是一个time.Duration，本质是一个int64，add()中传入的是一个time.Duration，所以直接用原本设置好的初始值相加后传入就行了
// 	var testchange time.Duration = 3600000000000 * 10 //自定义单位
// 	changedtime := now.Add(testchange)
// 	fmt.Printf("changedTime2: %v\n", changedtime)
// 	fmt.Printf("相差时间为%v\n", now.Sub(changedtime))                      //外减内
// 	fmt.Printf("两个时间是否相同: %v\n", now.Equal(changedtime))               //不同时区也能正确比较
// 	fmt.Printf("now时间是否在changedtime之前: %v\n", now.Before(changedtime)) //now时间是否在changedtime之前，相等返回false
// 	fmt.Printf("now时间是否在changedtime之后: %v\n", now.After(changedtime))  //now时间是否在changedtime之后，相等返回false

// }
// func timeFormat() {
// 	/* 	const (
// 		Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
// 		ANSIC       = "Mon Jan _2 15:04:05 2006"
// 		UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
// 		RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
// 		RFC822      = "02 Jan 06 15:04 MST"
// 		RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
// 		RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
// 		RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
// 		RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
// 		RFC3339     = "2006-01-02T15:04:05Z07:00"
// 		RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
// 		Kitchen     = "3:04PM"
// 		// Handy time stamps.
// 		Stamp      = "Jan _2 15:04:05"
// 		StampMilli = "Jan _2 15:04:05.000"
// 		StampMicro = "Jan _2 15:04:05.000000"
// 		StampNano  = "Jan _2 15:04:05.000000000"
// 	) */
// 	now := time.Now()
// 	fmt.Printf("now: %v\n", now)
// 	fmt.Printf("now.Format(time.ANSIC): %v\n", now.Format(time.ANSIC))
// 	// 自定义格式
// 	// 06和2006代表年(表示年份的后两位、年份的全四位)
// 	// 01和Jan代表月(数字月份和英文简写的月份)
// 	// 02和2代表日(用0补齐、不用0补齐)
// 	// MON、Monday代表星期几
// 	// T15、15、03代表小时(分别表现形式为T[hour]、24小时制、12小时制)
// 	// 04和4代表分钟(用0补齐、不用0补齐)
// 	// 05和5代表秒(用0补齐、不用0补齐)
// 	// 000、000000、000000000毫秒、微秒、纳秒
// 	// PM上午还是下午
// 	// 7表示数字时区(-7和Z7表示带正负，-0700)
// 	// MST表示时区
// 	fmt.Printf("now.Format(\"2006-01-02 MON 03:04:05 PM z7 MST\"): %v\n", now.Format("2006-01-02 MON 03:04:05 PM Z07 MST"))

// 	loc, _ := time.LoadLocation("Asia/Shanghai")
// 	timeobj, _ := time.ParseInLocation("2006/01/02 15:04:05", "2009/11/20 08:45:03", loc)
// 	/* 通过参数1的格式，解析参数2的时间字符串，根据参数3的位置信息，返回一个time.Time类型的结构*/
// 	fmt.Printf("timeobj: %v\n", timeobj)
// 	fmt.Printf("timeobj.Sub(now): %v\n", timeobj.Sub(now))
// }
// func main() {
// 	/*
// 		time包提供了测量和显示时间的功能

// 		type Time struct {
// 				wall uint64
// 				ext  int64
// 				loc *Location
// 		}

// 		type Duration int64
// 		const (
// 			minDuration Duration = -1 << 63
// 			maxDuration Duration = 1<<63 - 1
// 		)

// 		type Month int
// 		const (
// 			January Month = 1 + iota
// 			February
// 			March
// 			April
// 			May
// 			June
// 			July
// 			August
// 			September
// 			October
// 			November
// 			December
// 		)

// 		type Weekday int
// 		const (
// 			Sunday Weekday = iota
// 			Monday
// 			Tuesday
// 			Wednesday
// 			Thursday
// 			Friday
// 			Saturday
// 		)
// 	*/
// 	showtime()
// 	timestamp()
// 	timeFunction()
// 	timeFormat()
// }
