package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

/*
func Racer(url1, url2 string) (winner string) {
	startA := time.Now()
	http.Get(url1)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(url2)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return url1
	}

	return url2
}*/

/*// select 允许同时在多个channel等待。第一个发送值的channel胜出，case中的代码会被执行
func Racer(a, b string, timeout time.Duration) (winner string, err error) {
	// 如果a或b谁先有结果，谁先返回，但如果测试达到10秒，那么time.After会发送一个信号并返回一个erro
	select {
	case <-ping(a):
		return a, nil

	case <-ping(b):
		return b, nil

	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}*/
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error)  {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error)  {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waitting for %s and %s", a, b)
	}
}


func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		// 这个案例并不关心channel中发送的类型，只是想发送一个信号来说明已经发送完了，所以返回bool就可以了
		ch <- true
	}()

	return ch
}