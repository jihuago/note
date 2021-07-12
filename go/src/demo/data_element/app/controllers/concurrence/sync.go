package concurrence

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	m = make(map[string]int)
	x1 int64 = 0
	wgG = sync.WaitGroup{}
)

// sync.Map
// GO语音中内置的map不是并发安全的
func SyncMapDemo()  {
	//syncMap()
	//SyncMapDemo()
	noLockDemo()
}

// Go语音的sync包中提供了一个开箱即用的并发安全版map-sync.Map
// 开胸即用表示不用像内置的map一样使用make函数初始化就能直接使用
func syncMap()  {
	var m = sync.Map{}

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func get(key string) int {
	return m[key]
}

func set(key string, value int)  {
	m[key] = value
}

// 并非并发安全的
func mapDemo()  {

	wg := sync.WaitGroup{}

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 不加锁demo
// 下面两个goroutine在访问和修改x1变量的时候存在数据竞争，导致最后的结果与期待的不符
func noLockDemo()  {
	wgG.Add(2)
	go func() {
		for i := 0; i < 150000; i++ {
			x1 = x1 + 1
		}
		wgG.Done()
	}()
	go func() {
		for i := 0; i < 15000; i++ {
			x1 = x1 + 1
		}
		wgG.Done()
	}()
	wgG.Wait()
	fmt.Println(x1)
}
