package controllers

import "fmt"

// 给任意类型，包括基本类型，但不包括指针类型添加方法，如果是基本类型，需要借助type关键字对类型进行再定义
type Integer int

func (a Integer) Equal(b Integer) bool {
	return a == b
}

func RunRqual()  {
	var a Integer = 2
	bool := a.Equal(3)
	fmt.Println(bool)
}

// 组合实现类的继承和方法重写
type Animal struct {
	name string
}

func (a Animal) Call() string {
	return "动物叫声"
}

type Dog struct {
	Animal
}
