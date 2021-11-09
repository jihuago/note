package main

import (
	"fmt"
	"time"
)

// channel的应用：1. 实现超时控制
func main() {
	counter := make(chan int)

	// 生产
	go func() {
		i := 0
		for true {
			counter <- i
			i++
		}
	}()

	go runing(counter)

	select {
	case <- time.After(10 * time.Millisecond):
		fmt.Println("时间到")
		close(counter)
	default:
		fmt.Println("default")
	}

}

// 实现超时控制
// 10ns，计算一下可以打印多少数字
func runing(counter chan int)  {

	number, ok := <- counter
	fmt.Println(ok)
	for ok {
		fmt.Println("n:", number)
	}
}
