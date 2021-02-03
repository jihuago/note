package goroutime

import (
	"fmt"
	"time"
)

/*
	实现并行的for循环
		下面for循环的每一个迭代是并行完成的

	信号量模式
		下面的片段阐明：协程通过在通道ch中放置一个值来处理结束的信号。main协程等待<-ch直到从中获取到值。

 */


func TestGoFor()  {
	data := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}
	//fmt.Println(data)
	// 下面for循环的每一个迭代是并行完成的 在for循环中并行计算迭代可能带来很好的性能提升。
	for i, v := range data {
		go func(i , v int) {
			fmt.Println("key:", i, "value:", v)
		}(i, v)
	}

	time.Sleep(1e9)
}

func Test()  {
	data := []int{
		1, 2, 3,
	}

	ch := make(chan int)

	go compute(ch, data)

	res := <- ch

	fmt.Println(res)
}

func compute(ch chan int, data []int)  {
	// 协程通过在通道ch中放置一个值来处理结束的信号。
	ch <- sum1(data)

}

func sum1(data []int) int  {
	var result int = 0

	for _, v := range data {
		result += v
	}

	return result
}
