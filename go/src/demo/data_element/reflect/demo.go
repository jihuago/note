package reflect

import (
	"fmt"
	"reflect"
)

type Employer interface {
	get() string
}

type Employ struct {
	name string
}

func (e *Employ) get() string  {
	return e.name
}

// 在GO语言中，使用refect.TypeOf()函数可以获得任意值的类型对象，程序通过类型对象可以访问任意值的类型信息
func reflectType(e interface{})  {
	v := reflect.TypeOf(e)

	//res := v.get()

	fmt.Printf("type:%T, value:%v\n", v, v)
}

func TestReflect()  {
	var e Employer

	e1 := &Employ{"jack"}
	e = e1
	reflectType(e)

	i := 1
	reflectType(i)
}
