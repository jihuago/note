package controllers

import (
	"fmt"
	"strconv"
)

// strconv包实现了基本数据类型与其字符串表示的转换，主要有以下常用函数：Atoi()  Itia() parse系列 format系列  append系列
// string与int类型转换
//Atoi() 将字符串类型的整数转换为int类型
func DemoStr()  {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can not convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1)
	}

	// Itoa() 函数用于将int类型数据转换为对应的字符串表示
	// func Itoa(i int) string

	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("v:%v type: %T\n", s2, s2)

	// isPrint() 返回一个字符是否是可打印的，和unicode.Isprint一样，r必须是：字母、数字、标点、符号、ASCII空格
	// CanBackquote() 返回字符串s是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串

	// Format系列函数实现了将给定类型数据格式化为string类型数据的功能
	// func FormatBool(b bool) string   根据b的值返回"true"或"false"

	
}