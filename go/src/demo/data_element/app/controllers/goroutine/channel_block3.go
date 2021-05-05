package goroutine

import (
	"fmt"
	"time"
)

// 写一个通道证明它的阻塞性，开启一个协程接受通道的数据，持续15秒，然后给通道放入一个值。在不同阶段打印消息并观察输出

func RunChannelBlock3()  {
	ch := make(chan int)
	go logData(ch)

	time.Sleep(2e9)
	go func(ch chan int) {
		ch <- 133
	}(ch)

	time.Sleep(3e9)
}


// 跟进bug
func logData(ch chan int)  {

	for true {
		fmt.Println(<-ch)
	}
}


