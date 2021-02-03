package goroutime

import (
	"fmt"
	"time"
)

/*
	* 通过一个或多个通道交换数据进行协程同步
		* 通信是一种同步形式：通过通道，两个协程在通信中某刻同步交换数据。无缓冲通道成为了多个协程同步的完美工具。
		* 甚至可以在通道两端互相阻塞对方，形成了叫做死锁的状态。Go运行时会检查并panic，停止程序
		* 无缓冲通道会被阻塞。设计无阻塞的程序可以避免这种情况，或者使用带缓冲的通道
 */

func f1(in chan int) {
	fmt.Println(<-in)
}

// 解释为什么下边这个程序会导致 panic：所有的协程都休眠了 - 死锁！
// 解释： 运行时会检查所有的协程（本例只有一个）是否在等待什么，这意味着程序将无法继续执行
func Testf1()  {
	out := make(chan int)

/*	go func() {
		for  {
			out <- 2
		}
	}()*/

	out <- 2

	go f1(out)

	time.Sleep(1e9)
}
