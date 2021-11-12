package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
output:
9
0
1
2
3
4
5
6
7
8
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/home/vagrant/code/note/go/test/groutine/demo1.go:18 +0x96
exit status 2

原因：因为设置了只有一个P，所以for循环里面产生的goroutine都会进入P的runnext和本地队列，而不会涉及到全局队列
每次生成出来的goroutine都会第一时间塞到runnext，而i从1开始，runnetxt已经有goroutine在了，所以这时会把old goroutine移动到P的本地队列中去，
再把new goroutine放到runnext.
因此这后当一次 i 为 9 时，新 goroutine 被塞到 runnext，其余 goroutine 都在本地队列。
runnext 里的 goroutine 的执行优先级是最高的，因此会先打印出 9，接着再执行本地队列中的 goroutine 时，
按照先进先出的顺序打印：0, 1, 2, 3, 4, 5, 6, 7, 8


 */
func main() {
	var wg = sync.WaitGroup{}
	n := 10
	wg.Add(n)
	runtime.GOMAXPROCS(1)
	for i := 0; i < n; i++ {
		i := i
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
	//var ch = make(chan int)
	//<- ch
}