package structDemo

import "time"

/*
	1. Go方法是作用在接受者上的一个函数，接受者是某种类型的变量。因此方法是一种特殊类型的函数
	2. 接受者类型几乎可以是任何类型，不仅仅是结构体类型。可以是int、bool、string或数组的别名类型。但不能是一个接口类型
	3. 接受者不能是一个指针类型，但是它可以是任何其他允许类型的指针
    4. 在GO中，类型的代码和绑定在它上面的方法可以不放置在一起，它们可以存在于不同的源文件。但：类型代码、方法必须是同一个包
 */

// 将time.Time放在结构体中的匿名类型
type myTime struct {
	time.Time // 匿名字段
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}
