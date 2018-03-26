package main

import (
	"fmt"
	"time"
)

func main() {
	/*durationStr := "2018-03-26 22:26:00"                          //定时器休眠时间
	layout := "2006-01-02 15:04:05"                               //解析字符串为time.Time用的模板
	loc, _ := time.LoadLocation("Local")                          //获取loc,用于解析时间
	duration, _ := time.ParseInLocation(layout, durationStr, loc) //将时间字符串解析为time.Time
	fmt.Println(duration)
	durationUnix := duration.Unix()      //将duration转为unix时间戳
	triker := time.Tick(5 * time.Second) //定时器。设置5秒间隔时间
	for i := range triker {              //i是定时器返回的channel中的时间数据
		fmt.Println("do task", i)
		if i.Unix() >= durationUnix { //比较时间戳的大小，根据情况决定是否结束定时器的任务
			fmt.Println("fnish the task")
			break
		}
	}
	*/
	useTime()
}

//统计该函数执行时间，精确到微秒
func useTime() {
	startTime := time.Now().UnixNano() //当前时间的纳秒时间戳
	time.Sleep(time.Second * 5)
	endTime := time.Now().UnixNano()
	passTime := (endTime - startTime) / 1000
	fmt.Printf("cost %d us", passTime)
}
