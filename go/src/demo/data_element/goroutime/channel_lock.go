package goroutime

/*
	用带缓冲通道实现一个信号量
		信息量是实现互斥锁（排外锁）常见的同步机制，限制对资源的访问，解决读写的问题，比如没有实现信号量的sync的Go包，使用带缓冲的通道可以轻松实现：
			* 带缓冲通道的容量和要同步的资源

	给通道使用for循环
		for循环的range语句可以用在通道ch上，便可以从通道中获取值，像这样：
			for v := range ch {
				fmt.Printf("The value is %v\n", v)
			}
		它从指定通道中读取数据直到通道关闭，才继续执行下边的代码。很明显，另外一个协程必须写入ch（不然代码就阻塞在for循环了），
而且必须在写入完成后才关闭

	通道的方向
		通道类型可以用注解来表示它只发送或者只接收
			var send_only chan<- int

*/
