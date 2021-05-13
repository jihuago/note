package channel

import (
	"math/rand"
	"time"
)

/*type WorkerPoll struct {
	maxGoroutineNum int
	goroutineNum int
}*/

// 生成一个int64类型的随机数，发送到jobChan
func SeedInt64Random(ch chan<- int64) chan<- int64 {

	for true {
		int63 := rand.Int63()
		ch <- int63
		time.Sleep(1e8)
	}

	return ch
}

func readJobChan() {
	
}




