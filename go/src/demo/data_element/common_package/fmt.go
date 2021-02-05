package common_package

import (
	"fmt"
	"os"
	"time"
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

	定时器
		使用time.Tick来设置定时器，定时器的本质上是一个通道(channel)
		ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
		for i := range ticker {
			fmt.Println(i)
		}
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

// time.Time类型表示时间。通过time.Now()函数获取当前的时间对象
func TestTime()  {
	now := time.Now()
	year := now.Year()
	Month := now.Month()
	day := now.Day()

	unix := now.Unix() // 当前时间戳
	unixNano := now.UnixNano() // 当前纳秒时间戳

	fmt.Printf("Year:%d, day:%02d, Month:%02d,时间戳:%v,纳秒：%v\n", year, day, Month, unix, unixNano)

	later := setTimeAfter(time.Minute * 3)

	// 时间类型自带的方法Format进行格式化，Go语言中格式化时间模板不是常见的Y-m-d H:M:S 而是 2006年1月2号15点04分（记忆口诀2006 1 2 3 4）
	fmt.Println(later, now.Format("2006/01/02"))

	// 解析字符串格式的时间
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(loc)

	time := Date(time.Now(), "2006/01/02 03:04:05")
	fmt.Println(time)
}

func setTimeAfter(duration time.Duration) (time.Time) {
	now := time.Now()
	later := now.Add(duration)

	return later
}

func DemoTicker()  {
	ticker := time.Tick(time.Second)

	for i := range ticker {
		fmt.Println(i) // 每秒都会执行
	}
}

/*

 */
func Date(time time.Time, format string) string {

	return time.Format(format)
}
