package functions

import "fmt"

// 定义函数类型
// 定义函数类型和其他类型类似，同时后半部分和匿名函数类似，只不过没有函数实现。
type UserName func(name string) string

func DemoFunction() func() {
	var uname UserName = func(name string) string {
		return fmt.Sprintf("name:%v", name)
	}

	res := uname("jack")
	//fmt.Println(res)

	// 函数作为返回值：在Go中，这样的函数一定是匿名函数
	return func() {
		fmt.Println(res)
	}

}


