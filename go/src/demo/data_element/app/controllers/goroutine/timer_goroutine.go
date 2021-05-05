package goroutine

import (
	"fmt"
	"time"
)

// time.Tick(d) 以d为周期给返回的通道发送时间，d是纳秒数。当你想返回一个通道而不必关闭它的时候这个函数非常有用。

func RunTimer()  {
	//rate_per_sec := 10
	var dur time.Duration = 2e9
	chRate := time.Tick(dur) // chRate阻塞了更高的频率。每秒处理的频率可以根据机器负载资源的情况而增加或减少

	requests := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
	}

	// time.Tick(d) 每隔d向通道发送时间，<-chRate然后取出
	for req := range requests {
		<- chRate
		fmt.Println(req)
	}
}
