package goroutine

import (
	"fmt"
	"time"
)


// 限制请求频率
// 扩展上边的代码，思考如何承载周期请求数的暴增（提示：使用带缓冲通道和计时器对象）。
func LimitRequest()  {

	// 每秒3个请求，最大同时4个请求
	duration := time.Second
	//concurrencyNum := 3
	concurrencyNumMax := 4

	ch := make(chan int, concurrencyNumMax)
	ticker := time.NewTicker(duration)

	defer ticker.Stop()

	requests := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
	}

	for key, req := range requests {
		ch <- key
		<- ticker.C

		wg.Add(1)
		go work(ch, req)

		fmt.Println(req)
	}

	wg.Wait()
}

func work(ch chan int, k int)  {
	defer wg.Done()
	fmt.Println("doing work:", k)
	<-ch
}