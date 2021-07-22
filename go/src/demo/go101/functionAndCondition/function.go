package functionAndCondition

import "fmt"

func ConditionDemo()  {
	ifDemo()
}

// GO里switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch，但fallthrough强制执行后面的case代码
func ifDemo()  {

	// 在GO的if中声明的变量，这个变量的作用域只能在改条件内，其他地方就不起作用
/*	if i := test(); i > 2 {
		fmt.Println("big")
	}

	// //这个地方如果这样调用就编译出错了，因为i是条件里面的变量
	fmt.Println(i)*/

	// wrong
/*	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)*/
}

// 当传一个参数值到被调用函数里时，实际上是传了这个值的一份copy
// 传指针使得多个函数能操作同一个对象
// 传指针比较轻量级(8byte)，只是传内存地址，我们可以永指针传递体积大的结构体。
// GO语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。（若函数需）
func test() (i int) {
	return 1
}

func fortest()  {
	for i, j := 1, 2; i < 10; i++ {
		fmt.Println(i, j)
	}
}
