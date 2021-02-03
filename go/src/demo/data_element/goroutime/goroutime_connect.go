package goroutime

import (
	"fmt"
	"time"
)

// 协程间的通信
/*
	* GO有一种特殊的类型，通道(channel)，就像一个可以用于发型类型化数据的管道，由其负责协程之间的通信，从而避开所有由共享内存导致的陷阱
	* 通过通道进行通信的方式保证了同步性。
	* 数据在通道中进行传递：在任何给定时间，一个数据被设计为只有一个协程可以对其访问，所以不会发生数据竞争。属鸡的所有权（可以读写数据的能力）也因此被传递
	* 声明通道的格式
		var identifier chan datatype
		未初始化的通道的值是nil
	* 通道只能传输一种类型的数据，比如 chan int 或 chan string，所有的类型都可以用于通道，空接口interface也可以。甚至可以创建通道的通道
	* 通道实际上是类型化信息的队列：使数据得以传输。它是先进先出的结构所以可以保证发送给他们的元素的顺序。
		通道也是引用类型，所以我们使用make()函数来给它分配内存。
		var ch1 chan string
		ch1 = make(chan string)
		或  ch1 := make(chan string)
		// 构建一个int通道的通道
		chan := make(chan int)

		// 函数通道
		funcChan := make(chan func())

	* 通信操作符 <-
		* 信息按照箭头的方向流动
		* 流向通道（发送）
			ch <- int1 表示：用通道ch发送变量int1
		* 从通道流出（接收），三种方式
			1. int2 = <- ch  表示：变量int2从通道接收数据，变量int2如果没有声明： int2 := <- ch

			2. <- ch 可以单独调用获取通道的下一个值，当前值会被丢弃，但是可以用来验证
				if <- ch != 1000 {}
			同一个操作符 <- 既用于发送也用于接收，但Go会根据操作对象弄明白该干什么。
 */

func Testgoroutime()  {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string)  {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string)  {
	var input string
	//time.Sleep(2e9)
	for {
		input = <- ch
		fmt.Printf("%s ", input)
	}
}
