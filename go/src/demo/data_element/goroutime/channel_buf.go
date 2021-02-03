package goroutime

import (
	"fmt"
	"time"
)

/*
	同步通道：使用带缓冲的通道
		* 一个无缓冲通道只能包含1个元素，有时显得很局限。我们给通道提供了一个缓存，可以在扩展的make命令中设置它的容量
			buf := 100
			// buf 是通道可以同时容纳的元素个数
			ch1 := make(chan string, buf)
		* 在缓冲满载之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，直到缓冲空了。
		* 缓冲容量和类型无关，所以可以（尽管可能导致危险）给一些通道设置不同的容量，只要他们拥有同样的元素类型。内置的cap函数可以返回缓冲区的容量
		* 如果容量大于0，通道就是异步的了：缓冲满载或变空之前通信不会阻塞，元素会按照发送的顺序被接收。
			如果容量是0或者未设置，通信仅在收发双方准备好的情况下才可以成功
		* 同步  ch := make(chan type, value)
			value == 0 阻塞，同步
			value > 0 非阻塞（取决于value元素），异步
		若使用通道的缓冲，你的程序会在“请求”激增的时候表现更好：更具伸缩性。在设计算法时首先考虑使用无缓冲通道，只在不确定的情况下使用缓冲
*/

var sumResult int = 0

// 使用带缓冲的通道
func TestChannelBuf() {
	data := []int{
		1, 2, 3, 4, 5,
	}

	ch := make(chan int)

	// 放入需要计算的整数
	go insertNumber(ch, data)

	go sum(ch)

	time.Sleep(1e9)
	fmt.Println("sumResult:", sumResult)

}

func insertNumber(ch chan int, data []int)  {

	for _, v := range data {
		ch <- v
	}

}

func sum(ch chan int)  {

	//var i int = 0

	for {
		//i = <-ch
		//sumResult += i

		sumResult += <-ch
	}

}

// 协程中用通道输出结果
/*
	为了知道计算何时完成，可以通过信道回报。例如：
		ch := make(chan int)
		go sum(bigArray, ch)

		sum := <- ch
 */
