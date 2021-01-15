package common

import (
	"fmt"
	"reflect"
)

// 指针
/*
	变量的本质是对一块内存空间的命名，可以通过引用变量名来使用这块内存空间存储的值，而指针的含义则指向存储这些变量值的内存地址。
	Go语言引入指针类型，主要基于两点考虑：1. 为程序员提供操作变量对应内存数据结构的能力 2. 为了提高程序的性能（指针可以直接指向某个变量值的内存地址
可以极大节省内存空间，操作效率也更高）

	指针在Go语言中有两个使用场景：
		1. 类型指针
		2. 切片

	作为类型指针，允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算
	切片比原始指针具备更强大的特性，而且更为安全。切片在发生越界时，运行时会报出宕机，并打出堆栈，而原始指针只会崩溃。

	* 操作符作为右值时，意义是取指针的值，作为左值时，也就是放赋值操作符的左边时，表示指针指向的变量。
归纳：*操作符的根本意义就是操作指针指向的变量

	注意事项
		1. 不能获取字面量或者常量的地址
			const i = 5
			ptr := &i // error
			ptr := &19 // error
		2. Go不支持指针运算
			c := p++ // p 为指针

*/

func TestPointer()  {
	// 定义一个指针类型
	a := 100
	var ptr *int // 指针类型
	ptr = &a
	fmt.Println(ptr) // ptr本身是一个内存地址
	fmt.Println(*ptr) // *ptr 获取指针指向内存地址存储的值

	// My Test
	arr := []int{1, 2, 3}
	var pter *[]int
	pter = &arr
	fmt.Println(pter)
	fmt.Println(*pter)

	dict := map[string]int {
		"name": 1,
		"sex": 2,
	}
	var pter1 *map[string]int
	pter1 = &dict
	fmt.Println(pter1)

	var number1 int = 2
	fmt.Println("number1地址：", &number1)

	//doSwap()

	//a, b := 1, 2
	//swap2(&a, &b)
	//fmt.Println(a, b)

	test1()
}

func test1() {
	var a int = 10
	var p *int = &a

	a = 100
	fmt.Println("a = ", a) // 100

	*p = 250
	fmt.Println("a = ", a) // 250
	fmt.Println("*p = ", *p) // 250

	a = 1000
	fmt.Println("*p = ", *p) // 1000
}

/*
	a => 1     a的地址 => abc  => 1
	b => 2     b的地址  => egf => 2

	b => abc

 */

func swap2(a, b *int)  {
	fmt.Println(*a, *b)
	*a, *b = *b, *a

	fmt.Println(*a, *b)
}

// 使用指针修改值  交换函数
func swap(a, b *int)  {

	// 取a 指针的值
	t := *a

	fmt.Println("t:", reflect.TypeOf(t))

	// 取b指针的值，赋给a指针指向的变量
	*a = *b

	// 将a指针的值赋给b指针指向的变量
	*b = t
}

func doSwap()  {
	x, y := 1, 2
	swap(&x, &y)

	var n1 int = 1
	ptr := &n1
	res := *ptr + 1

	fmt.Println(res)

	// 对一个空指针的反向引用时不合法的
	//var p *int = nil
	//*p = 1

	fmt.Println("x=", x, "y=", y)

}
