package concurrency


// 信道
/*
	ci := make(chan int) //整数无缓冲信道
	cs := make(chan string) // 字符串无缓冲信道
	cs := make(chan *os.File, 100) // 指向文件的指针的缓存信道

   无缓冲信道在通信时会同步交换数据，它能确保两个Go协程的计算处于确定状态

 */
