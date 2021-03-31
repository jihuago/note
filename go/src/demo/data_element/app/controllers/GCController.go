package controllers

import "fmt"

func DemoGC()  {
	gcTest()
}

/*
	垃圾回收是编程语言中提供的自动的内存管理机制，自动释放不需要的对象，让出存储器资源，无需程序员手动执行。

	Golang中的垃圾回收主要应用三色标记法，GC过程和其他用户goroutine可并行运行，但需要一定时间的STW(stop the world),STW的过程中，CPU
不执行用户代码，全部用于垃圾回收。

	如何理解go语言中的interface？记住以下三点：
		1. interface是方法声明的集合
		2. 任何类型的对象实现了在interface接口中声明的全部方法，则说明该类型实现了该接口
		3. interface可以作为一种数据类型，实现了该接口的任何对象都可以给对应的接口类型变量赋值

*/
func gcTest() {
	var phone Phone

	// 体现了多态的特性，同一个phone的抽象接口，分别指向不同的实体对象，调用的call()方法，打印的效果不同，就是体现了多态的特性
	phone = new(NokiaPhone)
	phone.call()

	phone = new(ApplePhone)
	phone.call()
}

type Phone interface {
	call()
}
type NokiaPhone struct {
}
type ApplePhone struct {
}

func (nokiaphone NokiaPhone) call() {
	fmt.Println("nokia")
}

func (iphone ApplePhone) call() {
	fmt.Println("iphone")
}

// 面向对象中的开闭原则：一个软件实体比如类、模块和函数应该对扩展开发，对修改关闭





