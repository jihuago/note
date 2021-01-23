package common

import (
	"errors"
	"fmt"
)

/*
	* Go语言的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用等，这些运行时错误会引起宕机
	* 有时宕机也是一种合理的止损方法，如果在损失发生时，程序没有因为宕机而停止，那么用户将会付出更大的代价
*/

// 手动触发宕机
// Go语言可以在程序中手动触发宕机，让程序崩溃，这样开发者可以及时发现错误，同时减少可能的损失
// Go语言程序在宕机时，会将堆栈和goroutine信息输出到控制台，所以宕机也可以知晓发生错误的位置
func Close()  {

	panic("crash")
}

// 怎样施加panic的保护措施，避免程序奔溃
// Go语言的内建函数recover专用于恢复panic。recover函数无需任何参数，并且会返回一个空接口类型的值。
func TestRecover()  {
	fmt.Printf("%s\n", "this is a test")

	// 尽量把defer语句写在函数体开始处，因为在引发panic语句之后的所有语句，都不会有任何执行机会。
	defer func () {
		fmt.Println("Enter defer function.")

		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}

		fmt.Println("Exit defer function.")
	}()

	// 引发panic
	panic(errors.New("something wrong"))
}

// 某些场景下，使用panic是一个很好的流程控制工具
func first()  {
	//panic(errors.New("first crash"))
	panic("first crash")
}

func second()  {
	panic(errors.New("second crash"))
}

func Do() (err error) {
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println(r)
			err := fmt.Errorf("Error: %v", r)
			fmt.Println(err)
		}

	}()

	first()
	second()

	return err
}