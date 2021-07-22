package _defer

import "fmt"

/*
	* defer底层代码实现
		func A() {
			r = deferproc(...) // 1. 先注册
			// code to do somethine
			runtime.deferreturn() // 2. 执行
			return
		}

	* defer链表
		defer 信息会注册到一个链表，而当前执行的goroutine持有这个链表的头指针，每个goroutine在运行时都有一个对应的结构体g，其中有一个_defer *_defer
字段指向defer链表头
		defer链表练起来的是一个个_defer结构体，新注册的defer会添加到链表头，执行时也是从头开始。所以defer才会表现为倒序执行。


	* defer语句的执行时间
		return 不是原子操作，执行过程是：保存返回值 => 执行defer => 执行ret跳转
		return x 底层实现
			返回值  = x
			运行defer
			RET 指令
	* defer数据结构
		type _defer struct {
			siz int32 // 参数和结果的内存大小
			started bool // defer是否已经执行
			openDefer bool
			sp uintptr // sp 记录的是注册这个defer的函数栈指针，通过sp可以判断自己注册的defer是否已经执行完
			pc uintptr // pc是deferproc的返回地址
			fn *funcval // defer关键字中传入的函数
			_panic *_panic // 触发延迟调用的结构体，可能为空
			link *_defer
		}

f: r = 100
main: i = 0, g= 100

 */

/*
	defer使用要点
		* 延迟对函数进行调用
		* 即时对函数的参数进行求值
		* 根据defer顺序反序调用
 */
var g = 100
func DeepDefer()  {
	//demo1()
	//demo2()
	fmt.Println(demo3())
}

func f() (r int) {
	r = g
	defer func() {
		r = 200
	}()

	fmt.Printf("f: r = %d\n", r)

	r = 0
	return r
}

func demo2()  {
	i := f()
	fmt.Printf("main: i = %d, g = %d\n", i, g)
}

func demo3() (ret int) {
	defer func() {
		ret ++
	}()
	return 0
}


func demo1()  {
	i := 1
	j := 2
	defer func(a int) {
		fmt.Printf("i = %d, j = %d \n", a, j)
	}(i)
	i = 10
	j = 20
	fmt.Println("OVER")

	// output:
	// OVER
	// i = 1
}
