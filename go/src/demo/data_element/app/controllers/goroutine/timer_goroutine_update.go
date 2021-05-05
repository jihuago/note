package goroutine

import (
	"fmt"
	"time"
)

// 限制请求频率
// 扩展上边的代码，思考如何承载周期请求数的暴增（提示：使用带缓冲通道和计时器对象）。
func LimitRequest()  {

	//ch := make(chan time.Time, 10)

	var dur time.Duration = 1e9
	ticker := time.NewTicker(dur)

	defer ticker.Stop()

	requests := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
	}

	for res := range requests {
		<- ticker.C
		fmt.Println("update:", res)
	}

}