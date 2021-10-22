package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	// context.Background()产生根Context，通过下面的方法衍生更多的子Context
	// 1. WithCancel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range generatorInt(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}

	// 从Done()方法中获取chan，如果能通过Done()方法收到值，意味着context已经发起了取消请求，我们
	//r := <- ctx.Done()
	//fmt.Println(r, ctx.Err())
	time.Sleep(time.Millisecond)
}

// 生成>=1的整数
func generatorInt(ctx context.Context) <-chan int {
	res := make(chan int)
	n := 1

	go func() {
		for {
			select {
			case r := <- ctx.Done():
				fmt.Println("test:", r, ctx.Err())
				return
			case res <- n:
				n++
			}
		}
	}()

	return res
}
