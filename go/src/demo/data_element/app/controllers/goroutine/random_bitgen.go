package goroutine

import "fmt"

func RandomBitgen()  {

	ch := make(chan int)

	//消费者
	go func() {
		for true {
			fmt.Print(<-ch, " ")
		}
	}()

	for true {
		select {
		case ch <- 0:
		case ch <- 1:

		}
	}
}
