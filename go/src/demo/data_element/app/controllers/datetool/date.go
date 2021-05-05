package datetool

import (
	"fmt"
	"time"
)

// Go语言日期和时间戳转换： 日期字符串 <=> time.Time <=> 时间戳
// 1. 字符串日期和时间戳不能直接转换，需要通过time.Time完成
// 2. 涉及字符串日期的，字符串日期格式一定要以Go诞生的时间为基准，不能是随意的时间

func Date2Time()  {
	fmt.Println(">> Date2Time")
	defer fmt.Println("<< Date2Time")

	// 一定要以Go诞生的时间为基准
	const dateFormat = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(dateFormat, "May 20, 2021 at 00:00am (UTC)")
	fmt.Println(t)

	// time.Parse() 将字符串时间转成time.Time
	const shortFormat = "2006-01-02 03:04"
	t1, _ := time.Parse(shortFormat, "2021-04-30 12:00")
	fmt.Println(t1)

	// 将time.Time转换成时间戳
	t1Unix := t1.Unix()
	fmt.Println("t1unix:", t1Unix)
	
}

// 时间格式化
func Time2Date()  {
	fmt.Println(">> Time2Date")
	defer fmt.Println("<< Time2Date")

	tm := time.Now()
	tmStr := tm.Format("2006/01/02 03:04:05")
	fmt.Println("tmStr:", tmStr)
}
