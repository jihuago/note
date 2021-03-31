package controllers

import (
	"log"
	"runtime"
	"time"
)

/*
	* 如何分析程序的运行时间与CPU利用率情况
		1. shell内置time指令
			$ time go run test.go

				real    0m3.066s 从程序开始到结束，实际度过的时间
				user    0m1.007s  程序在用户态度过的时间
				sys     0m0.862s 程序在内核态度过的时间

		2. /usr/bin/time指令
			这个指令比内置的time更详细一些，使用的时候需要使用绝对路径，加上参数-v
			$ /usr/bin/time -v go run test.go

		3. 内存占用情况查看
			$ top -p $(pidof ./snippet_mem)
				snippet_mem是go的可执行文件
*/
func DemoDebug() {
	//fmt.Println(2)
	runTest()
}

func test()  {
	container := make([]int, 8) // slice会动态扩容，用slice来做堆内存申请

	log.Println("===> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container =append(container, i)
	}
	log.Println("===> loop end.")
}

func runTest()  {
	log.Println("Start.")
	test()

	log.Println("force gc.")
	runtime.GC()

	log.Println("Done.")

	time.Sleep(3600 * time.Second)
}