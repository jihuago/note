package goroutine

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 无缓冲的channel，得有取，才能放得进去
func DemoChannelFirst()  {

	ch := make(chan int)

	wg.Add(1)
	// 运行这一个goroutine，等待着ch有人往里放东西
	go func() {
		defer wg.Done()
		fmt.Println("取ch：", <-ch)
	}()

	// 往ch里放1
	ch <- 1

	fmt.Println("Over!")
	wg.Wait()
}