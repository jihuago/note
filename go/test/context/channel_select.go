package main

import (
	"fmt"
)

// 例子： chan + select方式来通知goroutine结束
func main()  {

	var stop = make(chan struct{})

	for r := range readFile(stop) {
		fmt.Println(r)
		if r == 5 {
			stop <- struct{}{}
			break
		}
	}
}


func readFile(stop chan struct{}) <- chan int {
	result := make(chan int)
	n := 0

	go func() {
		for true {
			select {
			case <- stop :
				fmt.Println("Over!")
				close(stop)
				return
			case result <- n:
				n++
			}
		}
	}()

	return result
}
