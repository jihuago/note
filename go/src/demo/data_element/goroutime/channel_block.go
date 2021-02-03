package goroutime

import (
	"fmt"
	"time"
)

/*
		* 默认情况下，通信是同步且无缓冲的：在有接受者接受数据之前，发送不会结束。
		* 通道的发送/接受操作在对方准备好之前是阻塞的
			* 对于同一个通道，发送操作（协程或函数汇总的），在接收者准备好之前是阻塞的：如果ch中的数据无人接收，就无法再给通道传入其他数据；
			新的输入无法在通道非空的情况下传入。所以发送操作会等待ch再次变为可用状态：就是通道值被接收时。
			* 对于同一个通道，接收操作是阻塞的，直到发送者可用：如果通道中没有数据，接受者就阻塞了
*/

func TestChannle() {
/*	ch := make(chan int)
	go pump(ch)
	go seek(ch)

	time.Sleep(1e9)*/
	//fmt.Println(<-ch)

	// new
/*	ch := make(chan int)
	go read(ch)

	fmt.Println(<-ch)*/
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func seek(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

var arr []int

// 写一个通道证明它的阻塞性，开启一个协程接收通道的数据，持续15秒，然后给通道放入一个值
func read(ch chan int)  {
	time.Sleep(15e9)
	x := <- ch
	fmt.Println(x)
}