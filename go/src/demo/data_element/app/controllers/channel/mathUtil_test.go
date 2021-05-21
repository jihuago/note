package channel

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestSeedInt64Random(t *testing.T) {
	t.Run("test for SeedInt64Random", func(t *testing.T) {

		var jobChan = make(chan int64, 100)

		wg.Add(25)
		go SeedInt64Random(jobChan)



	})
}
