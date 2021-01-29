package common

import (
	"fmt"
	"os"
	"strings"
)

/*
	从命令行读取参数
		os包有一个string类型的切片变量os.Args，用来处理一些基本的命令行参数，它在程序启动后读取命令行输入的参数。
*/
func TestArgs() {
	who := "Alice "

	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}

	fmt.Println("Good Morning", who)
}

// 每个发型师的客单量（服务人数），购买会员数，回头率
// 10 - 12:30  2.5   13:30 - 17:30  4
func HelloWho() {
	names := "Hello "

	if len(os.Args) > 1 {
		names += strings.Join(os.Args[1:], " ")
	}
	fmt.Println(names)
}