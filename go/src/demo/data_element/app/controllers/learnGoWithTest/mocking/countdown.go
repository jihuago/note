package mocking

import (
	"fmt"
	"io"
	"time"
)

const write = "write"
const sleep = "sleep"

// 将time.Sleep抽离出来，方便再测试中控制它
// 用依赖注入的方式来代替真正的time.Sleep
type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep()  {
	time.Sleep(o.duration)
}

// 监听调用次数
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep()  {
	s.Calls++
}

// 用一种新的测试来检查操作的顺序是否正确
// CountdownOperatonsSpy同时实现了io.writer和Sleeper，把每一次调用记录到slice
// 在测试汇总，我们可以通过关心slice的顺序，就能知道函数的调用顺序
type CountdownOperationsSpy struct {
	Calls []string
}
func (s *CountdownOperationsSpy) Sleep()  {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error)  {
	s.Calls = append(s.Calls, write)
	return
}

// 需求： 从3开始依次往下，当到0时打印GO!并退出，要求每次打印从新的一行开始且打印间隔一秒的停顿
func Countdown(input io.Writer, sleeper Sleeper) {
	// Countdown函数作用就是将数据写到某处，io.writer就是作为Go的一个接口来抓取数据的一种方式

	for i := 3; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(input, i)
	}

	sleeper.Sleep()
	fmt.Fprint(input, "Go!")
}
