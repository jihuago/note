package concurrent

import (
	"context"
	"fmt"
	"time"
)

/*
	Context通常被译作上下文。一般理解为程序单元的一个运行状态、现场、快照
		每个Goroutine在执行之前，都要先知道程序当前的执行状态，通常将这些执行状态封装在一个Context变量中，传递给要执行的Goroutine中。
		上下文则几乎已经成为传递与请求同生存周期变量的标准方法。

		Context的创建和调用关系是层层递进的，也就是我们所说的链式调用，类似数据结构里的树，从根节点开始，每一次调用就衍生一个叶子节点。
			1. 首先生存根节点，使用context.Background方法生存，而后可以进行链式调用使用context包里的各类方法，
			2. Context的调用应该是链式的，通过WithCancel、WithDeadline，WithTimeout或WithValue派生出的新的Context，当父Context被取消时，其派生的所有Context都将取消
			3. 通过context.Withxxx都将返回新的Context和CancelFunc。调用CancelFunc将取消子代，移除父代对子代的引用，并且停止所有定时器。
			4. WithDeadline和WithTimeout对应的是timerCtx，两者是相似的，WithDeadline是设置具体的deadline时间，到达时间后，后代goroutine退出，而WithTimeout简单粗暴，
直接return WithDeadline(parent, time.Now().Add(timeout))
 */

type favContextKey string

func ContextDemo()  {
/*	wg := &sync.WaitGroup{}
	values := []string{"https://www.baidu.com/", "https://www.zhihu.com/"}
	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range values {
		wg.Add(1)
		context.WithValue(ctx, favContextKey("url"), url)
	}*/
	//contextDemo1()
	withDeadlineDemo()
}

func withDeadlineDemo()  {
	// 定义了一个50毫秒之后过期的deadline
/*	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)*/

	// 传递带有超时的上下文，告诉阻塞函数在超时结束后应该放弃其工作
	ctx, cancel := context.WithTimeout(context.Background(), 50 * time.Millisecond)

	defer cancel()

	// 然后使用一个select让主程序陷入等待：等待1秒后打印overslept退出或者等待ctx过期后退出。 因为ctx50秒后就过期，所以ctx.Done()会先接收到值，上面的代码会打印ctx.Err()取消原因。
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("test:", ctx.Err())
	}
	//fmt.Println("final")
}

func contextDemo1()  {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}

}