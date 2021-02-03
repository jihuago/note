package common_package

import (
	"fmt"
	"os"
)

/*
	fmt
		fmt包实现了类似C语言的printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分

	向外输出
		标准库fmt提供了以下几种输出相关函数
	Print系统函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾
添加一个换行符

	Fprint系统函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容
		func Fprint(w io.Writer, a ...interface{}) {n int, err error}

	Sprint系列函数会把传入的数据生成并返回一个字符串
		func Sprint(a ...interface{}) string

	Errorf 函数根据format参数生成格式化字符串并返回一个包含该字符串的错误
		func Errorf(format string, a ...interface{}) error
		通常使用这种方式来自定义错误类型，例如：
			err := fmt.Errorf("这是一个错误")
		Go1.13版本为fmt.Errorf函数新加了一个%w占位符用来生成一个可以包裹Error的Wrapping Error
			e := errors.New("原始错误e")
			w := fmt.Errorf("Wrap了一个错误%w", e)

	*printf 系列函数都支持fotmat格式化参数，在这里我们按照占位符将被替换的变量类型划分，方便查询和记忆。

		通用占位符
			%v     值的默认格式表示
			%+v    类似%v，但输出结构体时会添加字段名
			%#v    值的GO语法表示
			%T     打印值的类型
			%%     百分号

		布尔类型
			%t     true或false

		字符串和[]byte
			%s  直接输出字符串或者[]byte
			%x  每个字节用两字符十六进制数表示
		指针
			%p   表示为十六进制，并加上前导的ox
*/

func TestFmt()  {
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")

	s1 := fmt.Sprint("this is a test")

	age := 18
	name := "jack"
	s2 := fmt.Sprintf("name:%s, age:%d", name, age)

	fmt.Println(s1, s2)

	fmt.Printf("姓名：%+v \n", name)
	fmt.Printf("T:%T \n", name)
}
