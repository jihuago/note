package lock

import "sync"

// sync.RWMutex读写锁：允许多个只读操作并行执行，但写操作会完全互斥。这种锁叫做 多读单写锁


var muRw sync.RWMutex
var balanceRW int

func getbalance() int {
	muRw.RLock()
	defer muRw.RUnlock()
	return balanceRW
}
