package errHanlde

import (
	"errors"
	"fmt"
)

/*
	* Go没有try/catch异常机制，不能执行抛异常操作。但有一套defer-panic-and-recover机制
	* Go的设计者认为try/catch机制的使用太泛滥了，而且从底层向更高层级抛异常太耗费资源。
		他们给GO设计的机制也可以捕捉异常，但是更轻量，并且只应该作为（处理错误的）最后的手段
	* Go处理普通错误的方式
		1. 通过在函数和方法中返回错误对象作为他们的唯一或最后一个返回值
		2. 如果返回nil，则没有错误发生
		3. 主调函数总是应该检查收到的错误
	* 库函数通常必须返回某种错误提示给主调函数
	* panic and recover是用来处理真正的异常（无法预测的错误）而不是普通的错误
	* Go检查和报告错误条件的惯有方式
		1. 产生错误的函数会返回两个变量，一个值和一个错误码；如果后者是nil就是成功，非nil就是发生了错误
		2. 为了防止发生错误时正在执行的函数被终止，在调用函数后必须检查错误
*/

// 定义错误
var errNotFound error = errors.New("Not found error")

func TestDefineError()  {
	//fmt.Printf("error: %v", errNotFound)

	if res, err := Sqrt(-1) ; err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Println(res)
	}



}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}

	return f * f, nil
}