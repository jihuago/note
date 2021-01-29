package _interface

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
	Go语言不是一种“传统”的面向对象编程语言：它里面没有类和继承的概念
		* Go语言有非常灵活的接口概念，通过接口可以实现很多面向对象的特性。
		* 接口提供了一种方式来说明对象的行为
		* 接口定义了一组方法集，但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。
		* 接口定义格式
			// Namer是一个接口类型
			type Namer interface {
				Method1(param_list) return_type
				Method2(param_list) return_type
			}
		* 按照约定，只包含一个方法的 接口名字由方法名加[e]r后缀组成，例如Printer、Reader
			还有一些不常用的方式（当后缀er不合适），比如Recoverable，此时接口名以able结尾，或者以I开头
		* Go语言中的接口都很简短，通常它们会包含0个、最多3个方法
		* 一个类型可以实现多个接口
*/


type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32  {
	return sq.side * sq.side
}

func Init()  {
	s1 := &Square{1.2}

	var areaIntf Shaper
	areaIntf = s1

	fmt.Printf("The square has area:%f\n", areaIntf.Area())

}

// 可以定义一个具有此方法的接口valuable，定义一个使用valuable类型作为参数的函数showValue()，所有实现了valuable接口的类型都可以用这个方法
type stockPosition struct {
	ticker string
	sharePrice float32
	count float32
}

func (s stockPosition) getValue() float32  {
	return s.sharePrice * s.count
}

type car struct {
	make string
	model string
	price float32
}

func (c car) getValue() float32  {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable)  {
	fmt.Printf("value of the asset is %f\n", asset.getValue())
}

func Init2()  {
	var o valuable = stockPosition{"good", 577, 4}
	showValue(o)

	c := car{"BMW", "M3", 66500}
	showValue(c)

	var r io.Reader
	r = os.Stdin
	r = bufio.NewReader(r)
}

// 接口嵌套接口
// 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样
type ReadWrite interface {
	Read(b bytes.Buffer) bool
	Write(b bytes.Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

