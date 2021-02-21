package goroutime

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 启动多个goroutine
func RunManyGoroutine() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done() // goroutine结束就等级-1
			// math/rand实现了伪随机数生成器
			r := rand.New(rand.NewSource(time.Now().UnixNano())) // 有种子。通常以时钟，输入输出等特殊节点作为参数，初始化。该类型生成的随机数相比无种子时重复概率较低
																 // 无种子。可以理解为此时种子为1，Seek(1)。如果无种子编译后运行的结果是定值
			fmt.Println(r.Intn(100))
		}(i)
	}

	wg.Wait() // 等待所有登记的goroutine都结束
}
/*
	G => P队列
	P唤醒一个M，M寻找是否有空闲的P，如有则将该G对象移动到它本身。
*/