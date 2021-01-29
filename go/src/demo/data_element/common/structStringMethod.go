package common

import (
	"fmt"
	"strconv"
)

/*
	类型的String()方法和格式化描述符
		当定义了一个很多方法的类型时，常会使用String()方法来定制类型的字符串形式的输出，换句话说：一种可读性和打印性的输出。
		1. 如果类型定义了String()方法，它会被用在fmt.Printf()中生成默认的输出：等同于使用格式化描述符%v产生的输出。
		2. fmt.Print()和fmt.Println()也会自动使用String()方法

https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/10.7.md
*/

type Person1 struct {
	name string
	age int
}

func (p *Person1) Init()  {
	p1 := Person1{name: "梁一一", age: 1}

	fmt.Println(p1)
	fmt.Printf("%#v", p1)
}

func (p *Person1) String() string {
	return "(" + p.name + strconv.Itoa(p.age) + ")"
}
