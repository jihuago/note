package goroutine

import (
	"fmt"
	"time"
)

// 信号量是实现互斥锁常见的同步机制，限制对资源的访问，解决读写问题，比如没有实现信号量的sync的Go包，使用带缓冲的通道可以轻松实现：
// 1. 带缓冲通道的容量和要同步的资源容量相同
// 2. 通道的长度与当前资源被使用的数量相同
// 3. 容量减去通道的长度就是未处理的资源个数

type Empty interface {}
type semaphore chan Empty

func (s semaphore) P(n int)  {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) V(n int)  {
	for i := 0; i < 0; i++ {
		<- s
	}
}

func (s semaphore) Lock()  {
	s.P(1)
}

func (s semaphore) Unlock()  {
	s.V(1)
}

func RunLock()  {

	sum := 0

	go func(sum *int) {
		for i := 0; i < 3; i++ {
			*sum += i
		}
	}(&sum)

	//fmt.Println("sum:", sum)

	go func(sum *int) {
		for i := 0; i < 2; i++ {
			*sum -= i
		}
	}(&sum)

	time.Sleep(2e9)
	fmt.Println("sum:", sum)

}