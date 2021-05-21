package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

// 模拟程序从多个复制的数据库同时读取。只需要一个答案，需要一个答案，需要接受首先达到的答案
func Query()  {
	conns := []int{2, 3, 4, 5, 1}

	// ch必须是带缓冲的：以保证第一个发送进来的数据有地方可以方寸，确保放入首个数据总会成功。
	//ch := make(chan int, 1)
	ch := make(chan int, 1)

	for _, conn := range conns {
		// 正在执行的协程可以使用runtime.Goexit()停止
		go func(conn int) {
			select {
			case ch <- doQuery(conn):
			default:
			runtime.Goexit()

			}
		}(conn)

	}

	fmt.Println(<-ch)

}

func doQuery(conn int) int {
	t := time.Duration(conn)
	time.Sleep(t * time.Second)
	return conn
}
