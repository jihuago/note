package goroutine

import (
	"fmt"
	"time"
)

func f1(in chan int)  {
	fmt.Println(<-in)
}

func RunF1()  {
	out := make(chan int)

	go func() {
		out <- 2
	}()

	go f1(out)
	time.Sleep(1e9)
}
