package goroutine

import "fmt"

// 通道阻塞
//1. d对于同一个通道，发送操作，在接受者准备好之前是阻塞的：如果ch中的数据无人接受，就无法再给通道传入其他数据，新的输入无法在通道非空的情况下传入。
// 所以发送操作会等待ch再次变成可用状态
//2. 对于同一个通道，接受操作是阻塞的，直到发送者可用：如果通道中没有数据，接受者就阻塞了

func Get()  {
	ch := make(chan int)
	go pump(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func pump(ch chan int)  {
	for i := 0; ; i++ {
		fmt.Println("i:", i)
		ch <- i
	}
}
