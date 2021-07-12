package concurrence

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	wg sync.WaitGroup
	x int64
	l sync.Mutex
)

/*
原子操作：针对基本数据类型我们可以使用原子操作来保证并发安全，因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好。
互斥锁加锁操作因为涉及到内核态的上下文切换会比较耗时、代价比较高。
atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数实现同步更好。
 */

func AtomicDemo() {
	start := time.Now()

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//go atomicAdd() // 原子操作版add函数，是并发安全，性能优化加锁版
		go add()
	}
	wg.Wait()
	fmt.Println(x)
	end := time.Now()
	used := end.Sub(start)
	fmt.Println(used)
}

func add()  {
	x++
	wg.Done()
}

// 原子操作版加函数
func atomicAdd()  {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

// 互斥锁版加函数，是并发安全的，但是加锁性能开销大
func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	wg.Done()
}