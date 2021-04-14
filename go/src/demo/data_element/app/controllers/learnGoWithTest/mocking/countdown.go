package mocking

import (
	"fmt"
	"io"
)

// 需求： 从3开始依次往下，当到0时打印GO!并退出，要求每次打印从新的一行开始且打印间隔一秒的停顿
func Countdown(input io.Writer) {
	// Countdown函数作用就是将数据写到某处，io.writer就是作为Go的一个接口来抓取数据的一种方式

	fmt.Fprint(input, "3")
}
