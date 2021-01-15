package common

// 结构体
// 定义
type info struct {
	name string
	age int
}

/*

结构体定义的一般方式：
	type identifier struct {
		field1 type1
		field2 type2
	}
	结构体里的字段都有名字，如field1，如果字段在代码中从来也不会被用到，那么可以命名为_

	结构体的字段可以是任何类型，甚至是结构体本身，也可以是函数或者接口。

	赋值
		type T struct {a, b int}
		var s T
		s.a = 5
		s.b = 8

	使用new
		使用new函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：var t *T = new(T)

		惯用方法：t := new(T)
		声明var t T 也会给t分配内存，并零值化内存。t通常被称为类型T的一个实例或对象

 */


