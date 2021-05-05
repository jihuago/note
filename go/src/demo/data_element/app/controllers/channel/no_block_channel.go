package channel

import (
	"fmt"
)

func demoChannel1()  {

	names := make(chan int, 5)

	// 写入
	go func() {
		for i := 0; i < 10; i++ {
			names <- i
		}
	}()

/*	for j := 0; j < 10; j++ {
		fmt.Println(<-names)
	}*/

	go func() {
		for j := 0; j < 10; j++ {
			fmt.Println(<-names)
		}
	}()

	//time.Sleep(2 * time.Second)
	for  {
		
	}

}
