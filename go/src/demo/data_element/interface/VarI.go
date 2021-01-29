package _interface

import "fmt"

/*
	类型断言：如何检测和转换接口变量的类型
		一个接口类型的变量varI中可以包含任何类型的值，需要有一种方式来检测它的动态类型，即运行时在变量中存储的值的实际类型。
		* 通常使用类型断言来测试在某个时刻varI是否包含类型T的值
			V := varI.(T)
		* varI必须是一个接口变量，否则报错
		* 使用以下形式进行类型断言
			if v, ok := varI.(T); ok {
				Process(v)
				return
			}
			如果转换合法，v是varI转换到类型T的值，ok会是true;否则v是类型T的零值,ok是false
*/

type AbstractController interface {
	GetName() string
}

type Controller struct {
	name string
}

func (c *Controller) GetName() string  {
	return c.name
}

func CheckVarI()  {
	var c AbstractController

	contro := &Controller{"jack"}
	c = contro

	if v, ok := c.(*Controller); ok {
		fmt.Println("The type of c is : %T", v)
	}
}
