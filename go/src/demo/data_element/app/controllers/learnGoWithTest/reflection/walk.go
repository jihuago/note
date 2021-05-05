package reflection

import (
	"reflect"
)

// 函数walk(x interface{}, fu func(string)) 参数为结构体x，并对x中的所有字符串字段调用fn函数
// 完成上述，需要使用反射：计算中的反射提供了程序检查自身结构体的能力，特别是通过类型，这是元编程的一种形式。

// 为什么不通过将所有参数都定义为interface{}类型得到真正灵活的函数？
// 1. 作为函数的使用者，使用interface将失去对类型安全的检查。如果你想传入string类型的Foo.bar但是传入的是int类型的Foo.baz，编译器将
// 无法通知你这个错误。
// 2. 如果是interface{}，你必须检查传入的所有参数，并尝试断定参数的类型以及如何处理它们。这是通过反射实现。
// 总之，除非真的需要否则不要使用反射
// 如果先实现函数的多态性，请考虑是否可以围绕接口设计它，以便用户可以用多种类型来调用你的函数。

func walk(x interface{}, fn func(input string))  {
	val := reflect.ValueOf(x) // ValurOf该函数返回一个给定变量的Value，为我们提供了检查值的方法
	field := val.Field(0)

	// String() 以字符串的形式返回底层值，但如果这个字段不是字符串，程序就会报错
	fn(field.String())
}

