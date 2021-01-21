package common

import "fmt"

/*
	* 类型和作用在它上面定义的方法必须在同一个包里定义
		可以采用间接的方式：先定义该类型(如int或float)的别名类型，然后再为别名类型定义方法
	* 指针或值作为将诶守着
		* 鉴于性能的原因，recv最常见的是一个指向receiver_type的指针，特别是在receiver类型是结构体时，就更是如此时。
		* 方法将指针作为接受者不是必须的，如果 func (p Point3) Abs() float64 {} ，但这样做稍微有点昂贵，因为Point3是作为值传递给方法的
			因此传递的是它的拷贝。
	* 并发访问对象
		对象的字段（属性）不应该由2个或2个以上的不同线程在同一个时间去改变。如果在程序发生这种情况，为了安全并发访问，可以使用包sync中的方法。

	* 内嵌类型的方法和继承
		当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌
*/

type employee struct {
	name string
	salary float32
}

func (this *employee) giveRaise(percent float32) {
	this.salary += this.salary * percent

}

func TestA()  {

	p1 := employee{"jack", 12000}
	p1.giveRaise(0.02)

	fmt.Println(p1.salary)
}

// 如何在类型中嵌入功能
/*
	主要有两种方式来实现在类型中嵌入功能：
	1. 聚合或组合：包含一个所需功能类型的具名字段
	2. 内嵌：内嵌（匿名地）所需功能类型
*/
type Log struct {
	msg string
}

type Customer struct {
	Name string
	log *Log
}

func (l *Log) Add(s string)  {
	l.msg += "\n" + s
}