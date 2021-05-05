package goroutine

import (
	"fmt"
	"time"
)

func RunGoroutine2() {

	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string)  {
	ch <- "A"
	ch <- "B"
	ch <- "C"
	ch <- "E"
}

func getData(ch chan string)  {

	//time.Sleep(2e9)
	for true {
		fmt.Printf("%s \n", <-ch)
	}

}
