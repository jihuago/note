package common

import "fmt"

/*
	defer的执行原理
		1. defer和defer后的函数什么时候执行
		2. defer后函数里的变量值是什么时候计算的
	问题1：defer在defer语句处执行，defer的执行结果是把defer后的函数压入到栈，等待return或者函数panic后，再按先进后出的顺序执行被defer的函数
	问题2：defer的函数的参数是在执行defer时计算的，defer的函数中的变量的值是在函数执行时计算的

	defer及defer函数的执行顺序分2步：
		1. 执行defer，计算函数的入参的值，并传递给函数，但不执行函数，而是将函数压入栈
		2. 函数return语句后，或panic后，执行压入栈的函数，函数中变量的值，此时会被计算
 */
func DemodeferTrack()  {
	defertest1()
}

func defertest1() (x int)  {
	defer fmt.Printf("in defer: x = %d\n", x)
	x = 7
	return 9
}
