package string

import "fmt"

/*
	* Go语言字符串的数据结构
		type StringHeader struct {
			Data uintptr // 指向字节数组的指针
			Len int // 数组大小，字节数
		}
	* Go字符串作为只读的类型，并不会直接向字符串直接追加元素改变其本身的内存空间，所有在字符串上的写入操作都是通过拷贝实现的
	* 字符串是Go语言中的基础数据类型，字符串往往被看成一个整体，但它实际上是一片连续的内存空间，也可以将它理解成一个由字符组成的数组
	* Go语言中的字符串只是一个只读的字节数组
	字符串拼接
		* 字符串拼接会调用copy将输入的多个字符串拷贝到目标字符串所在的内存空间。新的字符串是一片新的内存空间，与原来的字符串也没有任何关联，一旦需要拼接的字符串非常大，拷贝带来的性能损失是无法忽略的
	字符串类型转换
		* 无论从哪种类型转换到另一种都需要拷贝数据，而内存拷贝的性能损耗会随着字符串和 []byte 长度的增长而增长。
 */

func DemoStr()  {
	str := "hello"
	//println([]byte(str))
	fmt.Println([]byte(str))
	fmt.Printf("%T \n", []byte(str))
	fmt.Printf("%T \n", str)

	var ints []int = make([]int, 2, 5)
	ints = append(ints, 3)
	fmt.Println(ints[2])
}