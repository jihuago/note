package funcs

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Download(url string) {
	fmt.Println("start download", url)
	time.Sleep(time.Second)
	wg.Done()
}

