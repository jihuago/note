package controllers

import "fmt"


func Test() {
	//fmt.Println("controllers\\test")
	demo("jack", 1, 2, 3, 4, 5)
}

// 可变参数：函数的参数数量不固定。Go中的可变参数通过在参数名后加...来标记
// 固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
// 本质上，函数的可变参数是通过切片来实现的
func demo(name string, x ...int)  {
	for _, v := range x {
		fmt.Println(v)
	}
}
