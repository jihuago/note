package common

import (
	"fmt"
	"runtime"
)

/*
	垃圾回收和SetFinalizer
		* Go开发者不需要写代码来释放程序中不再使用的变量和结构占用的内存。
		* 在Go运行时中有一个独立的进程，即垃圾收集器GC，会处理这些事情，GC搜索不再使用的变量然后释放它们的内存。可以通过runtime包访问GC进程
		* 通过调用runtime.GC()函数可以显式的触发GC，但这只在某些罕见的场景下才有用，比如当内存资源不足时调用runtime.GC()，它会在此函数执行
的点上立即释放一大片内存，此时程序可能会有短时的性能下降（因为GC进程在执行）
		* 如果需要一个对象obj被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：
			runtime.SetFinalizer(obj, func(obj *typeObj))
*/

func GetMemStatus()  {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("%d Kb\n", m.Alloc / 1024)
}