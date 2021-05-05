package goroutine

import (
	"fmt"
	"time"
)

func DemoTimer()  {

	tick := time.Tick(1e8)
	after := time.After(5e8)

	for true {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-after:
			fmt.Println("after.")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}

}
