package common

import "fmt"

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
	结构体的内存布局
		Go语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。

	递归结构体
		结构体类型可以通过引用自身来定义。这在定义链表或二叉树的元素时特别有用，此时节点包含指向临近节点的链接。

 */

// 定义一个链表： head su => data su => tail nil
// data字段用于存放有效数据，su指针指向后继节点
// 链表中的第一个元素叫head，它指向第二个元素；最后一个元素叫tail，它没有后继元素，所以su为nil值。
type Node struct {
	data float64
	su *Node
}

// 定义一个双向链表节点
type DoubleNode struct {
	prev *DoubleNode
	data []string
	next *DoubleNode
}

func Person(p *info)  {

	// 赋值
	//p1 := info{"jack", 23}

}

/*
	使用工厂方法创建结构体实例
		* Go语言不支持面向对象编程语言那样的构造子方法，但是可以会很容易在GO中实现。
			type File struct {
				fd int // 文件描述符
				name string // 文件名
			}
			// 下面是这个结构体类型对应的工厂方法，它返回一个指向结构体实例的指针
			func NewFile(fd int, name string) *File {
				if fd < 0 {
					return nil
				}
				return &File{fd, name}
			}
			// 调用
			f := NewFile(10, "./test.txt")

	* 查看结构体类型T的一个实例占用了多少内存
		size := unsafe.Sizeof(T{})

	* 匿名字段和内嵌结构体
		结构体可以包含一个或多个匿名字段，即这些字段没有显示的名字，只有字段的类型是必须的，此时类型就是字段的名字。
		在一个结构体中对于每一种数据类型只能有一个匿名字段

*/

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int // 匿名字段
	innerS
}

func TestOuterS() {
	// 使用结构体字面量
	outer := outerS{6, 2.3, 2, innerS{1, 2}}

	fmt.Println(outer)

	type info struct {
		height float64
		int
		string
	}

	type person struct {
		info
		int
		height int
	}

	p1 := info{height: 175.2, int: 2, string: "aa"}
	fmt.Println(p1)

	p2 := person{info{1.2, 2, "aa"}, 3, 12}
	fmt.Println(p2.info.height)
}

// 命名冲突
/*
	当两个字段拥有相同的名字该怎么办
		1. 外层名字会覆盖内层名字（但两者的内存空间都保留），这提供了一种重载字段或方法的方式；
		2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误。
*/

// 方法
/*
	* 在GO语言中，结构体就像是
*/

