package goroutine

import (
	"fmt"
	"time"
)

// 使用select切换协程
// select 监听进入通道的数据，也可以是用通道发送值的时候
// 在任何一个case中执行break或return，select 就结束了
// select做的就是：选择处理列出的多个通信情况的一个

// go test -run TestRunSelectDemo
func RunSelectDemo()  {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go calu([]int{1, 2, 3}, ch1)
	go calu([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, ch3)
	go calu([]int{1, 2, 3, 4}, ch2)

	select {
	case v := <- ch1:
		fmt.Printf("Received on channel 1: %d\n", v)
	case v := <- ch2:
		fmt.Printf("Received on channel 2: %d\n", v)
	case v := <- ch3:
		fmt.Printf("Received on channel 3: %d\n", v)
	}

	time.Sleep(1e9)

}

func calu(arr []int, ch chan int) chan int {

	sum := 0
	for _, v := range arr {
		sum += v
	}

	ch <- sum

	return ch
}



