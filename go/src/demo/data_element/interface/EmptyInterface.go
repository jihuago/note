package _interface

import "fmt"

/*
	空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
	* 空接口类型的变量可以存储任意类型的变量
	* 只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口
*/

// 空接口作为函数的参数，使用空接口实现可以接受任意类型的函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value: %v \n", a, a)
}

func TestEmptyInterface()  {

	// 传递一个数组
	var arr [2]int
	arr = [2]int{
		1, 99,
	}
	show(arr)

	//传递一个结构体
	type info struct{
		name string
	}
	i := &info{name: "jack"}

	show(i)

}

