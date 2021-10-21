package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// WaitGroup控制并发的方式，可以控制多个goroutine同时完成
// WaitGroup适用于好多个goroutine协同做一件事情，因为每个goroutine做的都是这件事的一部分，只有全部的goroutine都完成，这件事才算完成
func main() {
	wg.Add(2)

	go func() {
		fmt.Println("任务1完成")
		time.Sleep(time.Second)
		wg.Done()
	}()

	go func() {
		fmt.Println("任务2完成")
		time.Sleep(time.Millisecond)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("所有任务完成")
}
