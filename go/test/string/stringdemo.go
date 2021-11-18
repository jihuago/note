package main

import (
	"fmt"
	"unicode"
)

// 字符串
// Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型一样。
// Go语言里的字符串的内部实现使用UTF-8编码

// Go语言的字符有以下两种：
// 1. uint8类型，或者叫byte，代表了ASCII码的一个字符
// 2. rune类型，代表一个UTF-8字符
func main() {
	//var str string = "hi"
	var str string = "hi啊，" // len()返回一个字符串中的字节数目。 一个中文在utf-8占用三个字节

	fmt.Println(len(str))

	// 字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的，字符串是由byte字节组成，所以字符串的长度是byte字节的长度
	// 修改字符串   先将其他转换成[]rune或[]byte，完成后再转换成string。无论哪种转换，都会重新分配内存，并复制字节数组
	str1 := "big"
	byteS1 := []byte(str1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	fmt.Println(HanCounter(str))

	// 字符串连接
	// + 连接适用于短小的、常量字符串（明确的，非变量），因为编译器会给我们优化
	// Join是比较统一的拼接，不太灵活
	// fmt和buffer基本上不推荐
	// builder 从性能和灵活性上，都是较好的选择
}

// HanCounter 统计字符串中有多少个汉字
func HanCounter(s string) int {
	count := 0
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}

	return count
}


