package structDemo

import (
	"fmt"
	"reflect"
)

/*
	结构体中的字段还有一个可选的标签（tag）：它是一个附属于字段的字符串，可以是文档或其他的重要标记。
	1. 标签的内容不可以在一般的编程中使用，只有包reflect能获取
 */

type TagType struct { // 带有标签的结构体
	name string "name 字段"
	sex int "性别"
}

func refTag(tt TagType, ix int)  {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)

	fmt.Printf("%v\n", ixField.Tag)
}

// 命名冲突
// 1. 外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式
// 2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误
type A struct {
	a int
}

type B struct {
	a, b int
}

type C struct {
	A
	B
	a int
}



