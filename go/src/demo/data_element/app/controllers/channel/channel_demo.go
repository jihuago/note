package channel

import "fmt"

func RunChannel1()  {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for true {
			value, ok := <-ch1

			if !ok {
				break;
			}

			ch2 <- value * value
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}
}

