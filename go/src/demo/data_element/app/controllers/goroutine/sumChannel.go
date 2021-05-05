package goroutine

import "fmt"

func sum(arr []int, ch chan int)  {

	var result int

	for _, value := range arr {
		result += value
	}

	ch <- result
}

// 协程通过在通道ch中放置一个值来处理结束的信号。 main协程等待 <- ch直到从中获取到值
func DoSum()  {
	ch := make(chan int)
	go sum([]int{
		1, 2, 3, 4,
	}, ch)

	sum := <- ch

	fmt.Println(sum)

}
