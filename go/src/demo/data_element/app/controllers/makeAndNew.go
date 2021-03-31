package controllers

import (
	"fmt"
	"sync"
)

func DemoAboutNew()  {

	/*
			new内置函数，只接受一个参数，这个参数是一个类型，分配好内存后，返回一个指向该类型内存地址的指针。
			同时把分配的内存置为零，也就是类型的零值

			func new(Type) *Type
	 */

	var i *int
	i = new(int)
	*i = 10

	fmt.Println(*i)

	lockUser()

}

type userInfo struct {
	lock sync.Mutex
	name string
	age int
}
// 体验new函数内存置零的好处
func lockUser() {
	u := new(userInfo) // 默认给u分配到内存全部为0

	u.lock.Lock() // 可以直接使用，因为lock是0，是开锁状态

	u.name = "张三"

	u.lock.Unlock()

	fmt.Println(u)

	// 示例中的userInfo类型中的lock字段不用初始化，直接可以用来用，不会有无效内存引用异常，因为它已经被零值了。
	// 这就是new, 它返回的永远是类型的指针，指向分配类型的内存地址

}

/*
	make也是用于内存分配的，但是和new不同。
	make只用于：
		* chan
		* map
		* slice
	的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

	func make(t Type, size ...IntegerType) Type

	# make与new的异同
		相同
			* 堆空间的分配
		不同
			make:只用于slice map以及channel的初始化，无可替代
			new : 用于类型内存分配（初始值为0），不常用
 */


