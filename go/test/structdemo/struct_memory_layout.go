package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	A int8
	B int8
	C int8
}

type Bar struct {
	x int32
	y *Foo
	z bool
}

/*
 	* Go在编译的时候会按照一定的规则自动进行内存对齐
	* 内存对齐可以减少CPU访问内存的次数，加大CPU访问内存的吞吐量，如果不进行内存对齐，可能就会增加CPU访问内存的次数
	* CPU访问内存，并不是逐个字节访问，而是以字为单位访问。比如64位CPU的字长为bbytes

	* 对齐保证
		unsafe.Alignof() Alignof函数获取一个变量的对齐系数

	* 存储结构体的起始地址是对齐边界的倍数，把起始地址看成0，结构体的每个成员在存储时都要把这个起始地址当作地址0。然后在用相对地址来决定自己该放哪里。
例如：
	type T struct {
		a int8  // 在64位系统， 对齐值 为  1byte
		b int64 // 8byte  存放b的时候，地址必须是8的倍数
		c int32 // 4byte
		d int16 // 2byte.
	}

 */

type Demo1 struct {
	age int8
	m struct{}
}

func main() {
	var b1 Bar
	fmt.Println(unsafe.Sizeof(b1))

	fmt.Println(unsafe.Alignof(b1.z)) // 1 表示此字段必须按1的倍数对齐

	var s string
	s = ""
	fmt.Println(unsafe.Sizeof(s))

	// 当空结构体类型作为结构体的最后一个字段时，如果有指向该字段的指针，那么就会返回该结构体之外的地址。为了避免内存泄漏会额外进行一次内存对齐
	var d1 Demo1
	fmt.Println(unsafe.Sizeof(d1))

	// 在Golang中，没有集合的类型，所以一般是把map当作集合来用
	// 使用空结构体作为Map的值来实现一个类似Set的数据结构
	set := make(map[int]struct{}, 2)
	set[1] = struct{}{}
	set[2] = struct{}{}

	fmt.Println(set)
}
