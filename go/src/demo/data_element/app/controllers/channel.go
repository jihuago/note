package controllers

import (
	"fmt"
	"sync"
	"time"
)

func DemoChannel() {
	//testReadChannel()
	testWaitGroup()
}

// 这个函数会报错。因为main在开辟完两个goroutine后，立刻关闭了ch，结果：panic: send on closed channel
func testReadChannel()  {

	/*
		Channel特性：
		* 给一个nil channel发送数据，造成永远阻塞
		* 给一个nil channel接受数据，造成永远阻塞
		* 给一个已经关闭的channel发送数据，引起panic
		* 从一个已经关闭的channel接受数据，如果缓冲区为空，则返回一个零值
		* 无缓冲的channel是同步的，而有缓冲的channel是非同步的

		口诀：空读写阻塞，写关闭异常，读关闭空零

	 */
	ch := make(chan int, 1000)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch

			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()

	//close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

const N = 10
var wg = &sync.WaitGroup{}

// WaitGroup与goroutine的竞速问题
func testWaitGroup()  {
	// 下面代码存在的问题：结果不唯一，代码存在风险，所有goroutine未必都能执行到
	for i := 0; i < N; i++ {
		wg.Add(1) // Add(1) 放goroutine外
		go func(i int) {
			//wg.Add(1) // 因为goroutine执行太快了，导致wg.Add(1)还没有执行 main函数就执行完毕了
			println(i)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
}
