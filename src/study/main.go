package main

import (
	"fmt"
	"time"
)

func main() {
	//返回当前时间
	now := time.Now()
	fmt.Println(now)
	// 当前时间戳
	fmt.Println(now.Unix())
	//纳米级时间戳
	fmt.Println(now.UnixNano())
	//时间戳小数部分 单位那秒
	fmt.Println(now.Nanosecond())
}
